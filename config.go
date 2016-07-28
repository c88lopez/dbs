package main

import (
	"encoding/json"
	"fmt"
	"os"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
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
	check(err)

	configFilePath = dir + "/.dbs/config"
}

func (c *Config) loadConfig() {
	setConfigFilePath()

	configFile, err := os.Open(configFilePath)
	check(err)
	defer configFile.Close()

	decoder := json.NewDecoder(configFile)

	check(decoder.Decode(&c))
}

func setDatabaseConfigInteractively() {
	fmt.Print("Configuring database parameters...\n")

	var username, database string

	fmt.Print("Username: ")
	fmt.Scanln(&username)

	fmt.Print("Password: ")

	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	check(err)

	password := string(bytePassword)

	fmt.Print("\nDatabase: ")
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
	check(err)
	defer configFile.Close()

	configFile.Truncate(0)
	configJson, err := json.Marshal(config)
	check(err)

	_, err = configFile.Write(configJson)
	check(err)
}
