package statesDiff

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/c88lopez/dbs/src/config"
	"github.com/c88lopez/dbs/src/entity"
)

// func Diff
func Diff() error {
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

func diffCurrentNext() (schemaChanges, error) {
	dir, err := config.GetMainFolderPath()
	if nil != err {
		return schemaChanges{}, err
	}

	/*
		get paths (we can put these as get in another place
	*/
	statesDir := dir + "/states"
	historyDir := statesDir + "/history"
	currentLink := statesDir + "/current"

	currentState, _ := os.Readlink(currentLink)

	f, err := os.Open(historyDir)
	if err != nil {
		return schemaChanges{}, err
	}

	/*
		searching the following state (we need to handle if there is no next state)
	*/
	var nextState string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), currentState) {
			scanner.Scan()
			nextState = scanner.Text()

			break
		}
	}

	/*
		decoding both current and next states
	*/
	file, err := os.Open(statesDir + "/" + currentState)
	if nil != err {
		return schemaChanges{}, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	currentSchema := new(entity.Schema)
	err = decoder.Decode(currentSchema)
	if err != nil {
		return schemaChanges{}, err
	}

	file, err = os.Open(statesDir + "/" + nextState)
	if nil != err {
		return schemaChanges{}, err
	}
	defer file.Close()

	decoder = json.NewDecoder(file)

	nextSchema := new(entity.Schema)
	err = decoder.Decode(nextSchema)
	if err != nil {
		return schemaChanges{}, err
	}

	/*
		here we check deeply for differences between the current schema and the next one
	*/
	var sc schemaChanges

	var checkingTableIndex int
	var checkingColumnIndex int
	for _, csTable := range currentSchema.Tables {
		checkingTableIndex = -1

		csTableStatus := tableStatus{}

		/*
			Looking if the table exists
		*/
		for nsTableIndex, nsTable := range nextSchema.Tables {
			if nsTable.Name == csTable.Name {
				checkingTableIndex = nsTableIndex
				break
			}
		}

		if checkingTableIndex == -1 {
			// Table does not exists => drop
			fmt.Printf("Should drop table %s", csTable.Name)
			csTableStatus.Status = "drop"
		} else {
			// Table found, checking definition and columns
			if csTable.DefaultCharset != nextSchema.Tables[checkingTableIndex].DefaultCharset ||
				csTable.Engine != nextSchema.Tables[checkingTableIndex].Engine {
				// The table has not the same definition
				fmt.Printf("Should alter table %s\n", csTable.Name)
				csTableStatus.Status = "alter"
			}

			for _, csColumn := range csTable.Columns {
				/*
					Looking if the column exists
				*/
				checkingColumnIndex = -1
				for nsColumnIndex, nsColumn := range nextSchema.Tables[checkingTableIndex].Columns {
					if nsColumn.Name == csColumn.Name {
						checkingColumnIndex = nsColumnIndex

						if csColumn.DataType != nsColumn.DataType ||
							csColumn.Nullable != nsColumn.Nullable ||
							csColumn.Key != nsColumn.Key ||
							csColumn.DefaultValue != nsColumn.DefaultValue ||
							csColumn.Extra != nsColumn.Extra {
							// The column has not the same definition
							fmt.Printf("Should alter column %s.%s\n", csTable.Name, csColumn.Name)

							if csColumn.DataType != nsColumn.DataType {
								fmt.Printf("\t- Type: %s => %s", csColumn.DataType, nsColumn.DataType)
								csTableStatus.ColumnsFinal = append(
									csTableStatus.ColumnsFinal, columnStatus{Status: "alter", ColumnFinal: nsColumn})
							}

							fmt.Println()
						}

						break
					}
				}

				if checkingColumnIndex == -1 {
					// Column does not exists => drop
					fmt.Printf("Should drop column %s.%s", csTable.Name, csColumn.Name)
					csTableStatus.ColumnsFinal = append(
						csTableStatus.ColumnsFinal, columnStatus{Status: "drop", ColumnFinal: csColumn})
				}
			}

			/*
				We should also checks table's FK and idxs...
			*/

		}

		if csTableStatus.Status != "" {
			sc.tables = append(sc.tables, csTableStatus)
		}
	}

	/*
		finally we should check for new tables
	*/
	for _, nsTable := range nextSchema.Tables {
		for _, csTable := range currentSchema.Tables {
			if nsTable.Name != csTable.Name {
				sc.tables = append(sc.tables, tableStatus{Status: "new", TableFinal: nsTable})

				break
			}
		}
	}

	return sc, nil
}

type schemaChanges struct {
	tables []tableStatus
}

type tableStatus struct {
	Status       string
	TableFinal   entity.Table
	ColumnsFinal []columnStatus
}

type columnStatus struct {
	Status      string
	ColumnFinal entity.Column
}
