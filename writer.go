package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type fileRecordWriter struct{}

func (frp *fileRecordWriter) Write(records Records, chunkSize int) error {
	err := loadChunks(records, chunkSize)
	if err != nil {
		return fmt.Errorf(" Error writing chunk: %v ", err)
	}
	return nil
}

func writeCsv(records Records, fileName string) error {
	file, err := os.Create(fileName)
	defer file.Close()
	if err != nil {
		return fmt.Errorf(" Error creating file: %v ", err)
	}
	writer := csv.NewWriter(file)
	defer writer.Flush()
	err = writer.Write(Header)
	if err != nil {
		return fmt.Errorf(" Error writing header: %v ", err)
	}

	for _, record := range records {
		err = writer.Write(record)
		if err != nil {
			return fmt.Errorf(" Error writing file: %v ", err)
		}
	}
	return nil
}

func loadChunks(records Records, chunkSize int) error {
	for i := 0; i < len(records); i++ {
		if (i+1)*chunkSize > len(records) {
			break
		}
		err := writeCsv(records[i*chunkSize:(i+1)*chunkSize], fmt.Sprintf("chunk_%d.csv", i+1))
		if err != nil {
			return fmt.Errorf(" Error writing csv number %d: %v ", i+1, err)
		}
	}
	return nil
}
