package handlers

import (
	"crypto/sha1"
	"os"
	"strings"
)

// GetLastState Will return the hash representing the last state in $PWD/.dbs/history
func GetLastState(historyFile *os.File) string {
	var err error
	lastLine := make([]byte, (sha1.Size*2)+1) // 40 chars (2 bytes each) + EOL ?

	for {
		_, err = historyFile.Read(lastLine)
		if nil != err {
			break
		}
	}

	return strings.TrimSpace(string(lastLine))
}

// GetCurrentState Will return the current schema state hash
func GetCurrentState() {

}
