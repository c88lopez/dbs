package config

import (
	"encoding/json"
	"fmt"
	"os"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

var configFilePath string
var Parameters = new(Config)

// Config struct
type Config struct {
	Driver   string `json:"driver"`
	Protocol string `json:"protocol"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

func SetConfigFilePath() error {
	dir, err := os.Getwd()
	if nil != err {
		return err
	}

	configFilePath = dir + "/.dbs/config"

	return err
}

func GetConfigTemplate() ([]byte, error) {
	baseConfig := Config{
		Driver:   "mysql",
		Protocol: "tcp",
		Host:     "localhost",
		Port:     "3306",
		Username: "dummy",
		Password: "dummy",
		Database: "dummy",
	}

	baseJson, err := json.Marshal(baseConfig)
	if nil != err {
		return []byte{}, err
	}

	return baseJson, nil
}

func LoadConfig() error {
	SetConfigFilePath()

	configFile, err := os.Open(configFilePath)
	if nil != err {
		return err
	}

	defer configFile.Close()

	decoder := json.NewDecoder(configFile)

	err = decoder.Decode(&Parameters)
	if nil != err {
		return err
	}

	return nil
}

func SetDatabaseConfigInteractively() error {
	fmt.Print("Configuring database parameters...\n")

	var protocol, host, port, username, database string

	fmt.Print("Protocol: ")
	fmt.Scanln(&protocol)

	fmt.Print("Host: ")
	fmt.Scanln(&host)

	fmt.Print("Port: ")
	fmt.Scanln(&port)

	fmt.Print("Username: ")
	fmt.Scanln(&username)

	fmt.Print("Password: ")

	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if nil != err {
		return err
	}

	password := string(bytePassword)

	fmt.Print("\nDatabase: ")
	fmt.Scanln(&database)

	LoadConfig()

	Parameters.Host = host
	Parameters.Port = port
	Parameters.Username = username
	Parameters.Password = string(password)
	Parameters.Database = database

	saveConfig(Parameters)

	return nil
}

func saveConfig(parameters *Config) error {
	SetConfigFilePath()

	configFile, err := os.OpenFile(configFilePath, os.O_WRONLY, 0600)
	if nil != err {
		return err
	}
	defer configFile.Close()

	configFile.Truncate(0)
	configJson, err := json.Marshal(parameters)
	if nil != err {
		return err
	}

	_, err = configFile.Write(configJson)
	if nil != err {
		return err
	}

	return nil
}
