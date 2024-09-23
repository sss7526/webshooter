package cli

import (
	"fmt"
	"os"
	"github.com/sss7526/goparse"
)

func ParseArgs() (map[string]interface{}) {
	parser := goparse.NewParser(
		goparse.WithName("Webshooter"),
		goparse.WithDescription("CLI utility to take screenshots and save PDFs of target web pages"),
	)

	parser.AddArgument("verbose", "v", "verbose", "Increase verbosity, shows http requests/responses and allowed/blocked status", "bool", false)
	parser.AddArgument("targets", "t", "targets", "Space separated list of one or more target URLs", "[]string", true)
	parser.AddArgument("pdf", "p", "pdf", "If specified, saves PDF copy of target webpage", "bool", false)
	parser.AddArgument("image", "i", "image", "If specified, saves screenshot of target webpage as a PNG", "bool", false)
	parser.AddArgument("translate", "T", "translate", "If specified, translates the target webpage before capture", "bool", false)

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
