package errors

import (
	"testing"
)

func TestHandleNoError(t *testing.T) {
	Handle(nil)

}
