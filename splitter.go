package main

import (
	"strings"
)

const Separator = "|"

type fileRecordSplitter struct{}

func (frs *fileRecordSplitter) Split(lines Lines) Records {
	var records Records
	for _, l := range lines {
		record := splitLine(l)
		if isValid(record) {
			records = append(records, record)
		}
	}
	return records
}

func splitLine(line string) Lines {
	return strings.Split(line, Separator)
}

func isValid(line Lines) (v bool) {
	var valid bool = true
	for _, v := range line {
		if v == "" {
			valid = false
			break
		}
	}
	return valid
}
