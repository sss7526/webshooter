package cli

import (
	"fmt"
	"os"
	"github.com/sss7526/goparse"
)

func ParseArgs() (map[string]interface{}) {
	parser := goparse.NewParser(
		goparse.WithName("\nWebshooter"),
		goparse.WithVersion("1.2.0\n"),
		goparse.WithDescription("CLI utility to take screenshots and save PDFs of target web pages.\nAlso can run queries against search engines and return the search result URLs.\nBlocks residual network requests to undesired URLs (trackers, ads, etc).\n"),
	)

	parser.AddArgument("verbose", "v", "verbose", "Increase verbosity, shows http requests/responses and allowed/blocked status", "bool", false)
	parser.AddArgument("targets", "t", "targets", "Space separated list of one or more target URLs", "[]string", false)
	parser.AddArgument("pdf", "p", "pdf", "If specified, saves PDF copy of target webpage", "bool", false)
	parser.AddArgument("image", "i", "image", "If specified, saves screenshot of target webpage as a PNG", "bool", false)
	parser.AddArgument("translate", "T", "translate", "If specified, translates the target webpage before capture", "bool", false)
	parser.AddArgument("tor", "P", "proxy", "If specified, connect to target over Tor (Tor must be installed and running)", "bool", false)
	parser.AddArgument("file", "f", "file", "Reads in target URLs from file. Cannot be used with --targets (-t) flag", "string", false)
	parser.AddArgument("query", "q", "query", "Sends query string to specified search engine", "[]string", false)
	parser.AddArgument("engine", "e", "engine", "search engine to query against (google and local instances of whoogle currently supported)", "string", false)

	parser.AddExclusiveGroup([]string{"targets", "file", "query"}, true)
	parser.AddExclusiveGroup([]string{"targets", "engine"}, false)

	parsedArgs, shouldExit, err := parser.Parse()
	if err != nil {
		fmt.Printf("Error parsing arguments: %v\n", err)
		if shouldExit {
			os.Exit(1)
		}
	}

	if shouldExit {
		os.Exit(0)
	}

	return parsedArgs
}
