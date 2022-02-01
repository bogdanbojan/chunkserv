package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var Header []string
var ChunkSize int

type fileRecordProvider struct{}

func (frp *fileRecordProvider) Get() (Lines, error) {
	chunkSize, err := getChunkSize()
	if err != nil {
		return nil, fmt.Errorf("Error getting chunk size: %v", err)
	}
	ChunkSize = chunkSize
	filePath, err := getFilePath()
	if err != nil {
		return nil, fmt.Errorf("Error getting file path: %v", err)
	}
	lines, err := getLines(filePath)
	if err != nil {
		return nil, fmt.Errorf("Error getting lines: %v", err)
	}
	return lines, nil
}

func getLines(filePath string) (lines Lines, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("Error opening file: %v", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	Header = getHeader(scanner)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return lines, nil
}

func getChunkSize() (cs int, err error) {
	fmt.Println("Set chunk size: ")
	var chunkSize int
	_, err = fmt.Scanln(&chunkSize)
	return chunkSize, err
}

func getFilePath() (fp string, err error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("(Leave empty for current dir) Set adress of file: ")
	filePath, err := reader.ReadString('\n')
	if filePath == "" || filePath == "\n" || filePath == "\r\n" {
		filePath = "input.txt"
	} else {
		filePath = strings.Replace(filePath, "\r\n", "", -1)
	}
	return filePath, err
}

func getHeader(scanner *bufio.Scanner) (h []string) {
	scanner.Scan()
	var header []string
	header = strings.Split(scanner.Text(), Separator)
	return header
}
