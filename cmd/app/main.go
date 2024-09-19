package main

import (
	"fmt"
	"os"
	"webshot/internal/processor"
	"webshot/internal/cli"
)

func main() {
	targets, err := cli.ParseArgs(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		cli.PrintHelp()
		os.Exit(1)
	}

	if len(targets) > 0 {
		processor.ProcessTargets(targets)
	}
}
