package main

import (
	"encoding/json"
	"log"
	"os"
)

// Config struct
type Config struct {
	Driver   string
	Username string
	Password string
	Database string
}

func (c *Config) loadConfig() *Config {
	configFile, err := os.Open("config.go")
	if err != nil {
		log.Panic(err)
	}

	decoder := json.NewDecoder(configFile)

	config := Config{}

	err = decoder.Decode(&config)
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
