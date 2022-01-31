package main

import (
	"strings"
)

type fileRecordSplitter struct{}

func (frs *fileRecordSplitter) Split(lines Lines) (r Records) {
	var records Records
	for _, l := range lines {
		record := splitLine(l)
		if isValid(record) {
			records = append(records, record)
		}
	}
	return records
}

func splitLine(line string) (l []string) { // Lines return
	return strings.Split(line, "|")
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
