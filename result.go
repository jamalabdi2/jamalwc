package main

import (
	"fmt"
)

type WcResult struct {
	NewLineCount int
	WordCounts   int
	ByteCounts   int
	Filename     string
	Filepath     string
}

func (r WcResult) FormatResults(opt Options) string {
	var result string

	if opt.Lines {
		result += fmt.Sprintf("%d ", r.NewLineCount)
	}
	if opt.Words {
		result += fmt.Sprintf("%d ", r.WordCounts)
	}

	if opt.Bytes {
		result += fmt.Sprintf("%d ", r.ByteCounts)
	}

	if r.Filename != "" {
		result += r.Filename
	}

	return result
}

func (wr *WcResult) PackStats(data []byte) {
	wr.NewLineCount, wr.WordCounts, wr.ByteCounts = ProcessStatistics(data)
}
