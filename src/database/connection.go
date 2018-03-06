package database

import (
	"database/sql"
	"fmt"

	"github.com/c88lopez/dbs/src/mainFolder"
	"github.com/fatih/color"

	_ "github.com/go-sql-driver/mysql"
)

var DbConnPool *sql.DB

func StartConnectionPool() error {
	fmt.Print("Initializing connection pool... ")

	if err := OpenConnectionPool(
		mainFolder.GetParameters().Username + ":" +
			mainFolder.GetParameters().Password + "@" + mainFolder.GetParameters().Protocol + "(" +
			mainFolder.GetParameters().Host + ":" + mainFolder.GetParameters().Port + ")/" +
			mainFolder.GetParameters().Database); err != nil {
		return err
	}

	color.Green("Done.\n")

	return nil
}

func OpenConnectionPool(dsn string) error {
	var err error

	if DbConnPool, err = sql.Open(mainFolder.GetParameters().Driver, dsn); nil != err {
		return err
	}

	if err = DbConnPool.Ping(); nil != err {
		return err
	}

	return nil
}

func CloseConnectionPool() error {
	return DbConnPool.Close()
}
