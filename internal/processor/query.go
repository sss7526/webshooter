package processor

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/cdproto/network"
	"time",
	"strings"
)

func Query(queries []string, engine string, verbose bool) {
	if engine != "" {
		for _, query := range queries {
			err := sendQuery(query, engine, verbose)
			if err != nil {
				fmt.Printf("Error sending query: %s\n", err)
			}
		}
	} else {
		fmt.Println("No search engine specified")
	}
}

func sendQuery(query, engine string, verbose bool) error {
	fmt.Printf("Query: %s\n", query)
	fmt.Printf("Engine: %s\n", engine)

	query = strings.Replace(query, " ", "+", -1)
	exclusionList := []string{"localhost", "search?", "maps?", "webhp?", "privacy?", "terms?", "cookies?", "websearch/answer"}

	var searchURL string
	var nextBtn string
	var nextBtnEval string

	jsScript := `Array.from(document.querySelectorAll('body a')).filter(a => {
		let parent = a.closest('nav, footer, header, [class*="nav"], [class*="header"], [class*="footer"]');
		return parent === null;
	})
	.map(a => a.href).filter(href => href && href.trim() !== "")`

	if engine == "google" {
		searchURL = fmt.Sprintf("https://www.google.com/search?q=%s", query)
		nextBtnEval = `document.querySelector('a#pnnext') !== null`
		nextBtn = `#pnnext`
	} else if engine == "whoogle" {
		searchURL = fmt.Sprintf("http://localhost:5000/search?q=%s", query)
		nextBtnEval = `document.querySelector('a[aria-label="Next page"]') !== null`
		nextBtn = `a[aria-label="Next page"]`
	} else {
		return fmt.Errorf("invalid engine: %s", engine)
	}

	keywordsToBlock := []string{"ads", "tracking", "analytics", "adservice", "counter", "track", "guestbook"}

	blockedURLS := []string{}
	for _, keyword := range keywordsToBlock {
		blockedURLS = append(blockedURLS, fmt.Sprintf("*%s*", keyword))
	}

	userAgent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebkit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"
	referrer := "https://www.google.com"

	opts := append(
		chromedp.DefaultExecAllocatorOptions[:],
		chromedp.NoDefaultBrowserCheck,
		chromedp.NoFirstRun,
		chromedp.UserAgent(userAgent),
		chromedp.Flag("disable-application-cache", true),
		chromedp.Flag("headless", true),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("no-sandbox", true),
		chromedp.Flag("ignore-certificate-errors", true),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	err := chromedp.Run(ctx,
		network.Enable(),
		network.SetBlockedURLS(blockedURLS),
	)
	if err != nil {
		return fmt.Errorf("failed to enable network events with blocked URLs: %w", err)
	}

	chromedp.ListenTarget(ctx, func(ev interface{}) {
		if verbose {
			switch ev := ev.(type) {
				
			case *network.EventRequestWillBeSent:
				shouldBlock := false
				badword := ""
				fmt.Printf("VALIDATING URL: %s\n\n", ev.Request.URL)
				for _, keyword := range keywordsToBlock {
					if strings.Contains(ev.Request.URL, keyword) {
						shouldBlock = true
						badword = keyword
						break
					}
				}

				if shouldBlock {
					fmt.Printf("BLOCKED Request: %s (contains '%s')\n\n", ev.Request.URL, badword)
				} else {
					fmt.Printf("ALLOWED Request URL: %s\n", ev.Request.URL)
					fmt.Printf("ALLOWED Request METHOD: %s\n", ev.Request.Method)
					fmt.Printf("ALLOWED Request HEADERS: %s\n\n", ev.Request.Headers)
				}

			case *network.EventResponseReceived:
				fmt.Printf("RESPONSE URL: %s\n", ev.Response.URL)
				fmt.Printf("RESPONSE STATUS: %d\n", ev.Response.Status)
				fmt.Printf("RESPONSE HEADERS: %s\n\n", ev.Response.Headers)
			}
		}
	})

	err = chromedp.Run(ctx,
		chromedp.ActionFunc(func(ctx context.Context) error {
			headers := make(map[string]interface{})
			headers["referer"] = referrer

			return network.SetExtraHTTPHeaders(network.Headers(headers)).Do(ctx)
		}),
		chromedp.Navigate(url),
		chromedp.WaitReady("body"),
		chromedp.Sleep(5*time.Second), // Lets images fully load first
	)

	if err != nil {
		return err
	}

	uniqueLinks := make(map[string]struct{})
	maxPages := 3
	for page := 0; page < maxPages; page++ {
		fmt.Printf("Processing Page %d...\n", page)

		var links []string
		var currentURL string
		err = chromedp.Run(ctx,
			chromedp.Evaluate(`document.querySelector('.jw8mI')?.remove(); document.querySelector('#KjcHPc')?.remove();`, nil), // Removes googles cookie acceptance splash page block
			chromedp.Evaluate(jsScript, &links),
			chromedp.Evaluate(`window.location.href`, &currentURL),
		)
		fmt.Printf("Navigated to page: %s\n", currentURL)

		if err != nil {
			fmt.Printf("Error evaluating javascript: %v", err)
		}

		links = removeURLsWithKeywords(links, exclusionList)
		deduplacateLinks(uniqueLinks, links)

		var nextBtnExists bool
		err = chromedp.Run(ctx,
			chromedp.Evaluate(nextBtnEval, &nextBtnExists),
		)

		if err != nil || !nextBtnExists {
			fmt.Println("No more pages to process, stopping...")
			break
		}

		err = chromedp.Run(ctx,
			chromedp.Click(nextBtn),
			chromedp.WaitReady("body"),
			chromedp.Sleep(5 * time.Second),
		)

		if err != nil {
			fmt.Printf("Error navigating to the next page: %v\n", err)
		}
	}

	for link := range uniqueLinks {
		fmt.Println(link)
	}

	return nil
}

func removeURLsWithKeywords(links, exclusionList []string) []string {
	var result []string
	for _, link := range links {
		if containsAny(link, exclusionList) {
			continue
		}
		result = append(result, link)
	}
	return result
}

func containsAny(link string, exclusionList []string) bool {
	for _, kw := range exclusionList {
		if strings.Contains(link, kw) {
			return true
		}
	}
	return false
}

func deduplicateLinks(uniqueLinks map[string]struct{}, links []string) {
	for _, link := range links {
		if _, exists := uniqueLinks[link]; !exists {
			uniqueLinks[link] = struct{}{}
		}
	}
}
