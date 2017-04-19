package statesDiff

import "os"

// func Diff
func Diff() error {
	var err error

	switch os.Args {
	default:
		err = diffCurrentPrevious()
	}

	return err
}

func diffCurrentPrevious() error {

	return nil
}
