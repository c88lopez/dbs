package database

import (
	"fmt"

	"github.com/c88lopez/dbs/src/config"
	"github.com/c88lopez/dbs/src/entity"
	"github.com/fatih/color"
)

func BuildSchemaState() (*entity.Schema, error) {
	fmt.Print("Building schema state... ")
	schema := new(entity.Schema)

	schema.Name = config.Parameters.Database
	err := schema.LoadInformationSchema(DbConnPool)
	if nil != err {
		return nil, err
	}

	if err := schema.FetchTables(DbConnPool); nil != err {
		return nil, err
	}

	color.Green("Done.\n")

	return schema, nil
}
