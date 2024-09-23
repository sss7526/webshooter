package main

import (
	"fmt"
	"os"
	"webshooter/internal/processor"
	"webshooter/internal/cli"
)

func main() {
	targets, verbose, err := cli.ParseArgs(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		cli.PrintHelp()
		os.Exit(1)
	}

	if len(targets) > 0 {
		processor.ProcessTargets(targets, verbose)
	}
}
