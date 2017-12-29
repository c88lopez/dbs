package statesDiff

import (
	"fmt"

	"github.com/c88lopez/dbs/src/entity"
)

func migrateNextStatus(sc schemaChanges) {

	for _, tableState := range sc.tables {
		if tableState.Status == "new" {
			createDefinitionsStatement := ""
			separator := ""

			for columnStateIndex, columnState := range tableState.TableFinal.Columns {
				if columnStateIndex > 0 {
					separator = ","
				}

				fmt.Printf("default: %#v", columnState.DefaultValue)

				createDefinitionsStatement += fmt.Sprintf(
					"%s %s %s %s %s %s %s",
					separator, columnState.Name, columnState.DataType, getNullableString(columnState),
					getDefaultString(columnState), getAutoIncrementString(columnState), getPrimaryKeyString(columnState))
			}

			fmt.Printf(
				"CREATE TABLE %s (%s)",
				tableState.TableFinal.Name, createDefinitionsStatement)
		}
	}
}

func getNullableString(cs entity.Column) string {
	nullableString := "NULL"

	if cs.Nullable == "NO" {
		nullableString = "NOT NULL"
	}

	return nullableString
}

func getDefaultString(cs entity.Column) string {
	defaultString := ""

	if cs.DefaultValue.Valid {
		defaultString = cs.DefaultValue.String
	}

	return defaultString
}

func getAutoIncrementString(cs entity.Column) string {
	autoIncrementString := ""

	if cs.Extra == "auto_increment" {
		autoIncrementString = "AUTO_INCREMENT"
	}

	return autoIncrementString
}

func getPrimaryKeyString(cs entity.Column) string {
	pkString := "PRIMARY KEY"

	if cs.Key == "" {
		pkString = ""
	}

	return pkString
}
