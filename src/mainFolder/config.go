package mainFolder

import (
	"encoding/json"
	"fmt"
	"os"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

var parameters = new(Config)

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

// GetConfigTemplate returns the default json
func GetConfigTemplate() ([]byte, error) {
	baseConfig := getDefaultParameters()

	baseJSON, err := json.Marshal(baseConfig)
	if nil != err {
		return nil, err
	}

	return baseJSON, nil
}

// LoadConfig get the config json and return it
func LoadConfig() error {
	file, err := os.Open(GetMainFolderPath())
	if nil != err {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	return decoder.Decode(&parameters)
}

// SetDatabaseConfigInteractively It will as the user for the config values
func SetDatabaseConfigInteractively() error {
	var err error

	defaultParameters := getDefaultParameters()

	parameters.Driver = getDriver(defaultParameters.Driver)
	parameters.Protocol = getProtocol(defaultParameters.Protocol)
	parameters.Host = getHost(defaultParameters.Host)
	parameters.Port = getPort(defaultParameters.Port)
	parameters.Username = getUsername(defaultParameters.Username)

	if parameters.Password, err = getPassword(); nil == err {
		parameters.Database = getDatabase(defaultParameters.Database)
		saveConfig(parameters)
	}

	return err
}

// GetParameters func
func GetParameters() *Config {
	return parameters
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
		Host:     "127.0.0.1",
		Port:     "3306",
		Username: "root",
		Password: "",
		Database: "dbs",
	}
}

func saveConfig(parameters *Config) error {
	file, err := os.OpenFile(GetConfigFilePath(), os.O_WRONLY, 0600)
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
