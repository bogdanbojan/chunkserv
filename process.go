package main

import "fmt"

type Records [][]string
type Lines []string

type RecordProvider interface {
	Get() (Lines, error)
}

type RecordSplitter interface {
	Split(Lines) Records
}

type RecordWriter interface {
	Write(Records, int) error
}

func process(rp RecordProvider, rs RecordSplitter, rw RecordWriter) error {
	records, err := rp.Get()
	if err != nil {
		return fmt.Errorf(" Error getting records: %s ", err)
	}

	splitRecords := rs.Split(records)

	err = rw.Write(splitRecords, ChunkSize)
	if err != nil {
		return fmt.Errorf(" Error writing records: %s ", err)
	}
	return nil
}
