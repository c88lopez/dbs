package statesDiff

import "os"

// func Diff
func CheckStatus() error {
	var err error
	var sc schemaChanges

	switch os.Args {
	default:
		sc, err = diffCurrentNext()
	}

	if err != nil {
		return err
	}

	migrateNextStatus(sc)

	return nil
}
