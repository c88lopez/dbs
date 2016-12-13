package test

import (
	"os"
	"testing"

	"github.com/c88lopez/dbs/command"
)

func TestGetFileLastLine(t *testing.T) {
	dummyFile, _ := os.Open("GetFileLastLine_dummyFile")
	command.GetFileLastLine(dummyFile)
}
