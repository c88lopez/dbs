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

	switch os.Args {
	default:
		err = diffCurrentPrevious()
	}

	return err
}

func diffCurrentPrevious() error {
	dir, err := config.GetMainFolderPath()
	if nil != err {
		return err
	}

	/**
	get paths (we can put these as get in another place
	*/
	statesDir := dir + "/states"
	historyDir := statesDir + "/history"
	currentLink := statesDir + "/current"

	currentState, _ := os.Readlink(currentLink)

	f, err := os.Open(historyDir)
	if err != nil {
		return err
	}

	/**
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

	/**
	decoding both current and next states
	*/
	file, err := os.Open(statesDir + "/" + currentState)
	if nil != err {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	currentSchema := new(entity.Schema)
	err = decoder.Decode(currentSchema)
	if err != nil {
		return err
	}

	file, err = os.Open(statesDir + "/" + nextState)
	if nil != err {
		return err
	}
	defer file.Close()

	decoder = json.NewDecoder(file)

	nextSchema := new(entity.Schema)
	err = decoder.Decode(nextSchema)
	if err != nil {
		return err
	}

	/**
	here we check deeply for differences
	*/
	var checkingTableIndex int
	var checkingColumnIndex int
	for _, csTable := range currentSchema.Tables {
		checkingTableIndex = -1

		/**
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
		} else {
			// Table found, checking definition and columns
			if csTable.DefaultCharset != nextSchema.Tables[checkingTableIndex].DefaultCharset ||
				csTable.Engine != nextSchema.Tables[checkingTableIndex].Engine {
				// The table has not the same definition
				fmt.Printf("Should alter table %s\n", csTable.Name)
			}

			for _, csColumn := range csTable.Columns {
				/**
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
							}

							fmt.Println()
						}

						break
					}
				}

				if checkingColumnIndex == -1 {
					// Column does not exists => drop
					fmt.Printf("Should drop column %s.%s", csTable.Name, csColumn.Name)
				}
			}

			/**
			We should also checks table's FK and idxs...
			*/

		}
	}

	return nil
}
