package processor

import (
	"context"
	"fmt"
	"webshooter/internal/validator"
	"log"
	"os"
	"strings"
	"time"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func ProcessTargets(targets []string, verbose bool) {
	fmt.Println("Processing targets")

	for _, target := range targets {
		target = validator.EnsureScheme(target)
		if validator.IsValidURL(target) {
			fmt.Printf("Valid URL: %s\n", target)

			filename := generateScreenshotFilename(target)

			err := takeScreenshot(target, filename, verbose)

			if err != nil {
				log.Printf("Error taking screenshot for %s: %s\n", target, err)
			} else {
				fmt.Printf("Screenshot saved: %s\n", filename)
			}
		} else {
			fmt.Printf("Invalid URL: %s\n", target)
		}
	}
}

func generateScreenshotFilename(url string) string {
	timestamp := time.Now().UTC().Format("20060102_150405")
	sanitizedURL := sanitizeFilename(url)
	filename := fmt.Sprintf("%s_%s_screenshot.png", sanitizedURL, timestamp)
	return filename
}

func sanitizeFilename(url string) string {
	replace := strings.NewReplacer("http://", "", "https://", "", "/", "_", ":", "", "?", "", "&", "", "=", "")
	return replacer.Replace(url)
}

func saveScreenshotToFile(filepath string, data []byte) error {
	err := os.MkdirAll("images", os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create 'images' directory: %v", err)
	}

	err = os.Writefile(filepath, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write screenshot to file: %v", err)
	}
	return nil
}

func savePDFToFile(filepath string, data []byte) error {
	err := os.MkdirAll("pdfs", os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create 'pdfs' direcotyr: %v", err)
	}
	filepath = strings.Replace(filepath, ".png", ".pdf", 1)
	err = os.WriteFile(filepath, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write PDF to file: %v", err)
	}
	return nil
}

func takeScreenshot(url, filename string, verbose bool) error {
	keywordsToBlock := []string{"ads", "tracking", "analytics", "adservice", "counter", "track", "guestbook"}
	
	blockedURLS := []string{}
	for _, keyword := range keywordsToBlock {
		blockedURLS = append(blockedURLS, fmt.Sprintf("*%s*", keyword))
	}

	userAgent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebkit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"
	referrer := "https://www.google.com"

	opts := append(
		chromedp.DefaultExecAllocatorOptions[:],
		chromedp.UserAgent(userAgent),
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
		return fmt.Error("failed to enable network events with blocked URLs: %w", err)
	}

	chromedp.ListenTargets(ctx, func(ev interfact{}) {
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

			case *networkEventResponseReceived:
				fmt.Printf("RESPONSE URL: %s\n", ev.Response.URL)
				fmt.Printf("RESPONSE STATUS: %s\n", ev.Response.Status)
				fmt.Printf("RESPONSE HEADERS: %s\n\n", ev.Response.Headers)
			}
		}
	})

	var buf []byte
	var pdfBuf []byte

	err = chromedp.Run(ctx,
		chromedp.ActionFunc(fun(ctx context.Context) error {
			headers := make(map[string]interface{})
			headers["Referer"] = referrer

			return network.SetExtraHTTPHeaders(network.Headers(headers)).Do(ctx)
		}),
		chromedp.Navigate(url),
		chromedp.WaitReady("body"),
		chromedp.Sleep(5 * time.Second), // Lets images fully load first
		chromedp.Evaluate(`document.querySelector('.jw8mI')?.remove(); document.querySelector('#KjcHPc)?.remove();`, nil) // Removes googles cookie acceptance splash page block
		chromedp.FullScreenshot(&buf, 100),
		chromedp.ActionFunc(func(ctx context.Context) error {
			pdf, _, err := page.PrintToPDF().WithPrintBackground(true).Do(ctx)
			if err != nil {
				return fmt.Errorf("failed to create PDF: %v", err)
			}
			pdfBuf = pdf
			return nil
		})
	)

	if err != nil {
		return err
	}

	filepath := fmt.Sprintf("images/%s", filename)

	err = saveScreenshotToFile(filepath, buf)
	if err != nil {
		return err
	}

	filepath = fmt.Sprintf("pdfs/%s", filename)
	err = savePDFToFile(filepath, pdfBuf)
	if err != nil {
		return err
	}

	return nil
}
