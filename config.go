package main

import (
	"encoding/json"
	"log"
	"os"
)

// Config struct
type Config struct {
	Driver   string `json:"driver"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

func (c *Config) loadConfig() *Config {
	configFile, err := os.Open("config.json")
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
