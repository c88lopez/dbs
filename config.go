package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/howeyc/gopass"
)

var configFilePath string
var config = new(Config)

// Config struct
type Config struct {
	Driver   string `json:"driver"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

func setConfigFilePath() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	configFilePath = dir + "/.dbs/config"
}

func (c *Config) loadConfig() {
	setConfigFilePath()

	configFile, err := os.Open(configFilePath)
	if err != nil {
		panic(err)
	}
	defer configFile.Close()

	decoder := json.NewDecoder(configFile)

	err = decoder.Decode(&c)
	if err != nil {
		panic(err)
	}
}

func setDatabaseConfigInteractively() {
	fmt.Print("Configuring database parameters...\n")

	var username, database string

	fmt.Print("Username: ")
	fmt.Scanln(&username)

	fmt.Print("Password: ")
	gopass.GetPasswdMasked()
	password, _ := gopass.GetPasswd()

	fmt.Print("Database: ")
	fmt.Scanln(&database)

	var config Config
	config.loadConfig()

	config.Username = username
	config.Password = string(password)
	config.Database = database

	saveConfig(config)
}

func saveConfig(config Config) {
	setConfigFilePath()

	configFile, err := os.OpenFile(configFilePath, os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer configFile.Close()

	configFile.Truncate(0)
	configJson, err := json.Marshal(config)
	if err != nil {
		panic(err)
	}

	_, err = configFile.Write(configJson)
	if err != nil {
		panic(err)
	}
}
