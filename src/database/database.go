package database

import (
	"fmt"

	"github.com/c88lopez/dbs/src/entity"
	"github.com/fatih/color"
)

// BuildSchemaState func
func BuildSchemaState() (*entity.Schema, error) {
	fmt.Print("Building schema state... ")
	schema := new(entity.Schema)

	err := schema.LoadInformationSchema(dbConnPool)
	if nil != err {
		return nil, err
	}

	if err := schema.FetchTables(dbConnPool); nil != err {
		return nil, err
	}

	color.Green("Done.\n")

	return schema, nil
}
