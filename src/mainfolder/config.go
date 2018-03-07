package mainfolder

import (
	"fmt"
	"syscall"

	"github.com/spf13/viper"
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

// SetDatabaseConfigInteractively It will as the user for the config values
func SetDatabaseConfigInteractively() error {
	getDriver()
	getProtocol()
	getHost()
	getPort()
	getUsername()
	getPassword()
	getDatabase()

	return viper.WriteConfig()
}

// GetParameters func
func GetParameters() *Config {
	return parameters
}

func getDriver() {
	var driver string

	fmt.Printf("Driver (%s): ", viper.Get("driver"))
	fmt.Scanln(&driver)

	if "" != driver {
		viper.Set("driver", driver)
	}
}

func getProtocol() {
	var protocol string

	fmt.Printf("Protocol (%s): ", viper.Get("protocol"))
	fmt.Scanln(&protocol)

	if "" != protocol {
		viper.Set("protocol", protocol)
	}
}

func getHost() {
	var host string

	fmt.Printf("Host (%s): ", viper.Get("host"))
	fmt.Scanln(&host)

	if "" != host {
		viper.Set("host", host)
	}
}

func getPort() {
	var port string

	fmt.Printf("Port (%d): ", viper.Get("port"))
	fmt.Scanln(&port)

	if "" != port {
		viper.Set("port", port)
	}
}

func getUsername() {
	var username string

	fmt.Printf("Username (%s): ", viper.Get("username"))
	fmt.Scanln(&username)

	if "" != username {
		viper.Set("username", username)
	}
}

func getPassword() {
	fmt.Print("Password: ")

	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if nil != err {
		panic(err)
	}

	viper.Set("Password", string(bytePassword))
}

func getDatabase() {
	var database string

	fmt.Printf("\nDatabase (%s): ", viper.Get("database"))
	fmt.Scanln(&database)

	if "" != database {
		viper.Set("database", database)
	}
}
