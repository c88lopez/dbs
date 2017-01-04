package command

import (
	"os"

	"github.com/c88lopez/dbs/src/common"
	"github.com/c88lopez/dbs/src/config"
	"github.com/c88lopez/dbs/src/database"
	"github.com/c88lopez/dbs/src/help"
)

func Dispatch() error {
	var err error

	if len(os.Args) == 1 {
		help.ShowHelp()
	} else {
		switch os.Args[1] {
		default:
			help.ShowHelp()
		case "init":
			err = common.GenerateMainFolder()
			break
		case "config":
			err = config.SetDatabaseConfigInteractively()
			break
		case "new":
			err = database.New()
			break
		}
	}

	return err
}
