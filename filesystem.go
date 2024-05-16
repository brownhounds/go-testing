package go_testing

import (
	"os"
	"testing"
)

func AssertFileDirExists(t *testing.T, filepath string) {
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		t.Fatal(err)
	}
}

func RemoveFileDir(t *testing.T, filepath string) {
	if err := os.RemoveAll(filepath); err != nil {
		t.Fatal(err)
	}
}
