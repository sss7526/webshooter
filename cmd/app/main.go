package main

import (
	"webshooter/internal/processor"
	"webshooter/internal/cli"
)

func main() {
	parsedArgs := cli.ParseArgs()

	targets := parsedArgs["targets"].([]string)

	saveToImage, ok := parsedArgs["image"].(bool)
	if !ok && !saveToImage {
		saveToImage = false
	} else {
		saveToImage = true
	}

	saveToPDF, ok := parsedArgs["pdf"].(bool)
	if !ok && !saveToPDF {
		saveToPDF = false
	} else {
		saveToPDF = true
	}

	verbose, ok := parsedArgs["verbose"].(bool)
	if !ok && !verbose {
		verbose = false
	} else {
		verbose = true
	}

	if len(targets) > 0 {
		processor.ProcessTargets(targets, verbose, saveToImage, saveToPDF)
	}
}
