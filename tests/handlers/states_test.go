package test

import (
	"os"
	"testing"

	"github.com/c88lopez/dbs/handlers"
)

func TestGetFileLastLine(t *testing.T) {
	lastLineValue := "d7fb97dafb90adf90b70a9df7b908adfb7adfb09"
	dummyFile, _ := os.Open("GetFileLastLine_dummyFile")

	if handlers.GetLastState(dummyFile) != lastLineValue {
		t.Fail()
	}
}
