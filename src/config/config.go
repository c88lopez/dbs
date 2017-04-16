package config

import (
	"encoding/json"
	"fmt"
	"os"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

var filePath string
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
	if nil == err {
		filePath = dir + "/.dbs/config"
	}

	return err
}

func GetConfigTemplate() ([]byte, error) {
	baseConfig := getDefaultParameters()

	baseJson, err := json.Marshal(baseConfig)
	if nil != err {
		return []byte{}, err
	}

	return baseJson, nil
}

func LoadConfig() error {
	SetConfigFilePath()

	file, err := os.Open(filePath)
	if nil != err {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	return decoder.Decode(&Parameters)
}

func SetDatabaseConfigInteractively() error {
	fmt.Print("Configuring database parameters...\n")

	var err error

	defaultParameters := getDefaultParameters()

	Parameters.Driver = getDriver(defaultParameters.Driver)
	Parameters.Protocol = getProtocol(defaultParameters.Protocol)
	Parameters.Host = getHost(defaultParameters.Host)
	Parameters.Port = getPort(defaultParameters.Port)
	Parameters.Username = getUsername(defaultParameters.Username)

	if Parameters.Password, err = getPassword(); nil == err {
		Parameters.Database = getDatabase(defaultParameters.Database)
		saveConfig(Parameters)
	}

	return err
}

func getDriver(defaultValue string) string {
	var driver string

	fmt.Printf("Driver (%s): ", defaultValue)
	fmt.Scanln(&driver)

	if "" == driver {
		driver = defaultValue
	}

	return driver
}

func getProtocol(defaultValue string) string {
	var protocol string

	fmt.Printf("Protocol (%s): ", defaultValue)
	fmt.Scanln(&protocol)

	if "" == protocol {
		protocol = defaultValue
	}

	return protocol
}

func getHost(defaultValue string) string {
	var host string

	fmt.Printf("Host (%s): ", defaultValue)
	fmt.Scanln(&host)

	if "" == host {
		host = defaultValue
	}

	return host
}

func getPort(defaultValue string) string {
	var port string

	fmt.Printf("Port (%s): ", defaultValue)
	fmt.Scanln(&port)

	if "" == port {
		port = defaultValue
	}

	return port
}

func getUsername(defaultValue string) string {
	var username string

	fmt.Printf("Username (%s): ", defaultValue)
	fmt.Scanln(&username)

	if "" == username {
		username = defaultValue
	}

	return username
}

func getPassword() (string, error) {
	fmt.Print("Password: ")

	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if nil != err {
		return "", err
	}

	return string(bytePassword), nil
}

func getDatabase(defaultValue string) string {
	var database string

	fmt.Print("\nDatabase(dbs): ")
	fmt.Scanln(&database)

	if "" == database {
		database = defaultValue
	}

	return database
}

func getDefaultParameters() Config {
	return Config{
		Driver:   "mysql",
		Protocol: "tcp",
		Host:     "172.17.0.2",
		Port:     "3306",
		Username: "root",
		Password: "",
		Database: "dbs",
	}
}

func saveConfig(parameters *Config) error {
	SetConfigFilePath()

	file, err := os.OpenFile(filePath, os.O_WRONLY, 0600)
	if nil != err {
		return err
	}
	defer file.Close()

	file.Truncate(0)
	jsonString, err := json.Marshal(parameters)
	if nil == err {
		_, err = file.Write(jsonString)
	}

	return err
}
