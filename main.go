package main

func main() {
	frp := fileRecordProvider{}
	frs := fileRecordSplitter{}
	frw := fileRecordWriter{}
	err := process(&frp, &frs, &frw)
	if err != nil {
		return
	}
}
