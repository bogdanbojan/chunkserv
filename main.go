package main

import "fmt"

func main() { // error?
	frp := fileRecordProvider{}
	frs := fileRecordSplitter{}
	frw := fileRecordWriter{}
	err := process(&frp, &frs, &frw)
	if err != nil {
		fmt.Println(err)
		return
	}
}
