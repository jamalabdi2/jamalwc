package main

import (
	"fmt"
	"os"
)

func main() {
	stat, err := os.Stdin.Stat()
	if err != nil {
		fmt.Printf("error getting stdin information: %v\n", err)
		return
	}
	mode := stat.Mode()

	opts, files, err := parseArgs(os.Args[1:])

	if err != nil {
		fmt.Println(err)
		return
	}

	switch {
	case mode.IsRegular():
		ProcessStdinAndDisplay(opts)
	case mode&os.ModeCharDevice != 0:
		HandleFile(opts, files)
	case mode&os.ModeNamedPipe != 0:
		fmt.Println("from pipe")
		ProcessStdinAndDisplay(opts)
	default:
		fmt.Println("Unknow device")
		return
	}
}
