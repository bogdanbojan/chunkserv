package main

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestGetChunkSize(t *testing.T) {

	content := []byte("100")
	tmpfile, err := ioutil.TempFile("", "")
	if err != nil {
		log.Fatal(err)
	}

	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal(err)
	}

	if _, err := tmpfile.Seek(0, 0); err != nil {
		log.Fatal(err)
	}

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }() // Restore original Stdin

	os.Stdin = tmpfile
	got, err := getChunkSize()
	if err != nil {
		t.Errorf("userInput failed: %v", err)
	}
	err = tmpfile.Close()
	if err != nil {
		log.Fatal(err)
	}

	want := 100

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestGetFilePath(t *testing.T) {}

func TestGetLines(t *testing.T) {}
