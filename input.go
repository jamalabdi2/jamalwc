package main

import (
	"fmt"
	"os"
)

func handleStdin() (*WcResult, error) {
	results := new(WcResult)

	byteData, err := ProcessReader(os.Stdin)

	if err != nil {
		return results, fmt.Errorf("error reading from stdin: %w ", err)
	}
	results.PackStats(byteData)
	return results, nil
}

func ProcessStdinAndDisplay(opts Options) {
	result, err := handleStdin()
	if err != nil {
		fmt.Println("Error " + err.Error())
		return
	}
	fmt.Println(result.FormatResults(opts))
}

func HandleFile(opts Options, files []string) {
	cwd, err := GetCwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, filename := range files {
		result, err := ProcessFile(filename, cwd)
		if err != nil {
			fmt.Println("Error processing file ", filename, err)
			return
		}
		line := result.FormatResults(opts)
		fmt.Println(line)
	}
}
