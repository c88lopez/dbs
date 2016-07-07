package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/fatih/color"
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

	configFilePath = dir + "/.dbsconfig.json"
}

func generateConfigFile() {
	fmt.Printf("Generating config file... ")

	baseConfig := Config{
		Driver:   "mysql",
		Username: "root",
		Password: "root",
		Database: "test",
	}

	json, err := json.Marshal(baseConfig)
	if err != nil {
		log.Panic(err)
	}

	setConfigFilePath()
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		err = ioutil.WriteFile(configFilePath, json, 0775)
		if err != nil {
			log.Panic(err)
		}

		color.Green("Done.\n")
	} else {
		color.Yellow("File already exist!.\n")
	}
}

func (c *Config) loadConfig() *Config {
	configFile, err := os.Open(configFilePath)
	if err != nil {
		log.Panic(err)
	}

	decoder := json.NewDecoder(configFile)

	err = decoder.Decode(&c)
	if err != nil {
		log.Panic(err)
	}

	return c
}

func (c *Config) getDriver() string {
	return c.Driver
}

func (c *Config) getUsername() string {
	return c.Username
}

func (c *Config) getPassword() string {
	return c.Password
}

func (c *Config) getDatabase() string {
	return c.Database
}
