package tests

import (
	"buildingHTTPServer/poker"
	"io"
	"testing"
)

func TestTapeWrite(t *testing.T) {
	file, clean := createTempFile(t, "12345")
	defer clean()
	tape := &poker.Tape{File: file}
	tape.Write([]byte("abc"))
	file.Seek(0, io.SeekStart)
	newFileContents, _ := io.ReadAll(file)
	got := string(newFileContents)
	want := "abc"
	assertEqual(t, got, want)
}
