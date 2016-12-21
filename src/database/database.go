package database

import (
	"database/sql"
	"fmt"

	"github.com/c88lopez/dbs/src/config"
	"github.com/c88lopez/dbs/src/entity"
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

func BuildSchemaState() (*entity.Schema, error) {
	fmt.Print("Building schema state... ")
	schema := new(entity.Schema)

	schema.Name = config.Parameters.Database
	err := schema.LoadInformationSchema(DbConnPool)
	if nil != err {
		return nil, err
	}
	schema.FetchTables(DbConnPool)

	color.Green("Done.\n")

	return schema, nil
}
