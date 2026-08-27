package main

import (
	"fmt"
	"strings"
)

type Options struct {
	Lines bool
	Bytes bool
	Words bool
}

func parseArgs(args []string) (Options, []string, error) {
	var option Options
	var files []string

	for _, arg := range args {
		switch arg {
		case "-l":
			option.Lines = true
		case "-c":
			option.Bytes = true
		case "-w", "-words":
			option.Words = true
		default:
			if strings.HasPrefix(arg, "-") {
				return option, nil, fmt.Errorf("unknown optio: %s", arg)
			}
			files = append(files, arg)
		}
	}
	if !option.Bytes && !option.Lines && !option.Words {
		option.Bytes = true
		option.Lines = true
		option.Words = true
	}
	return option, files, nil
}
