package cli

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func ParseArgs(args []string) ([]string, bool, error) {
	var tergets []string
	targetsSet := false
	var verbose bool = false

	if len(args) <= 1 {
		PrintHelp()
		os.Exit(0)
	}

	for i := 1; i < len(args); i++ {
		arg := args[i]

		if arg == "--help" || arg == "-h" {
			PrintHelp()
			os.Exit(0)
		} else if arg == "-v" || arg == "--verbose" {
			verbose = true

		}
		} else if arg == "--targets" || arg == "-t" {
			if targetsSet {
				return nil, verbose, errors.New("the --targets (-t) flag should only be specified once")
			}

			if i + 1 >= len(args) || strings.HasPrefix(args[i + 1], "-") {
				return nil, verbose, errors.New("no targets specified for --targets (-t)")
			}

			for i + 1 < len(args) && !strings.HasPrefix(args[i + 1], "-") {
				i++
				targets = append(targets, args[i])
			}

			targetsSet = true
		} else {
			return nil, verbose, errors.New("invalid argument: " + arg)
		}
	}

	if len(targets) == 0 {
		return nil, verbose, nil
	}

	return targets, verbose, nil
}

func PrintHelp() {
	helpMessage := `
Usage: goshot [OPTIONS]

Options:
    -h, --help        Show this help message and exit
    -t, --targets     Specify one or more targets. Example:

                      goshot --targets target1 target2 target3
`

	fmt.Println(helpMessage)
}
