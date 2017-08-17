package statesDiff

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

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
	dir, err := os.Getwd()
	if nil != err {
		return err
	}

	/**
	get paths (we can put these as get in another place
	*/
	statesDir := dir + "/.dbs/states"
	historyDir := statesDir + "/history"
	currentLink := statesDir + "/current"

	fmt.Printf("%s\n", currentLink)

	currentState, _ := os.Readlink(currentLink)

	fmt.Printf("%v\n", currentState)

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

	fmt.Printf("%v\n", currentSchema.Tables)

	fmt.Printf("nextState: %s\n", nextState)

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

	fmt.Printf("%v\n", nextSchema.Tables)

	/**
	here we check deeply for differences
	*/

	return nil
}
