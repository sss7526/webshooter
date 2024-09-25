package main

import (
	"bufio"
	"io"
	"os"
	"fmt"
	"path/filepath"
	"github.com/sss7526/webshooter/internal/processor"
	"github.com/sss7526/webshooter/internal/cli"
)

func main() {
	parsedArgs := cli.ParseArgs()

	targets, ok := parsedArgs["targets"].([]string)
	if !ok && len(targets) == 0 {
		targets = nil
	}

	query, ok := parsedArgs["query"].([]string)
	if !ok && len(query) == 0 {
		query = nil
	}

	saveToImage, ok := parsedArgs["image"].(bool)
	if !ok && !saveToImage {
		saveToImage = false
	}

	saveToPDF, ok := parsedArgs["pdf"].(bool)
	if !ok && !saveToPDF {
		saveToPDF = false
	}

	verbose, ok := parsedArgs["verbose"].(bool)
	if !ok && !verbose {
		verbose = false
	}

	translate, ok := parsedArgs["translate"].(bool)
	if !ok && !translate {
		translate = false
	}

	useTorProxy, ok := parsedArgs["tor"].(bool)
	if !ok && !useTorProxy {
		useTorProxy = false
	}

	engine, ok := parsedArgs["engine"].(string)

	fmt.Printf("Use Tor Proxy: %v\n", useTorProxy)

	filepath, ok := parsedArgs["file"].(string)
	if ok && filepath != "" {
		absPath, err := resolveFilePath(filepath)
		if err != nil{
			fmt.Printf("Error parsing file: %v\n", err)
			os.Exit(1)
		}
		file, err := os.Open(absPath)
		if err != nil {
			fmt.Printf("Error opening file: %v\n", err)
			os.Exit(1)
		}
		targets, err = readLines(file)
		if err != nil {
			fmt.Printf("%v", err)
			os.Exit(1)
		}
		defer file.Close()
	}

	if len(targets) > 0 {
		processor.ProcessTargets(targets, verbose, saveToImage, saveToPDF, translate, useTorProxy)
	} else if len(query) > 0 {
		processor.Query(query, engine, verbose)
	} else {
		fmt.Println("No targets or queries specified")
	}
}

func readLines(file *os.File) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	err := scanner.Err()
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}
	return lines, nil
}

func resolveFilePath(path string) (string, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	return absPath, nil
}

func ReadBlock(file *os.File) (string, error) {
	content, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("error reading file: %v", err)
	}
	return string(content), nil
}
