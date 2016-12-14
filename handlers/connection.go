package handlers

import (
	"database/sql"
	"fmt"

	"github.com/c88lopez/dbs/config"
	"github.com/fatih/color"
)

var DbConnPool *sql.DB

func StartConnectionPool() {
	fmt.Print("Initializing connection pool... ")

	OpenConnectionPool(
		config.Parameters.Username + ":" + config.Parameters.Password + "@" + config.Parameters.Protocol + "(" +
			config.Parameters.Host + ":" + config.Parameters.Port + ")/" + config.Parameters.Database)

	color.Green("Done.\n")
}

func OpenConnectionPool(dsn string) error {
	var err error

	DbConnPool, err = sql.Open(config.Parameters.Driver, dsn)
	if nil != err {
		return err
	}

	return nil
}

func CloseConnectionPool() error {
	return DbConnPool.Close()
}
