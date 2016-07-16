package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
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
		log.Panic(err)
	}

	configFilePath = dir + "/.dbs/config"
}

func (c *Config) loadConfig() *Config {
	setConfigFilePath()

	configFile, err := os.Open(configFilePath)
	if err != nil {
		log.Panic(err)
	}
	defer configFile.Close()

	decoder := json.NewDecoder(configFile)

	err = decoder.Decode(&c)
	if err != nil {
		log.Panic(err)
	}

	return c
}

func setDatabaseConfigInteractively() {
	fmt.Printf("Configuring database parameters...\n")

	var username, password, database string

	fmt.Printf("Username: ")
	fmt.Scanln(&username)

	fmt.Printf("Password: ")
	fmt.Scanln(&password)

	fmt.Printf("Database: ")
	fmt.Scanln(&database)

	config.loadConfig()
	config.Username = username
	config.Password = password
	config.Database = database

	configFile, err := os.OpenFile(configFilePath, os.O_WRONLY, 0600)
	if err != nil {
		log.Panic(err)
	}
	defer configFile.Close()

	configFile.Truncate(0)
	configJson, err := json.Marshal(config)
	if err != nil {
		log.Panic(err)
	}

	_, err = configFile.Write(configJson)
	if err != nil {
		log.Panic(err)
	}
}
