package database

import (
	"database/sql"
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/viper"

	// Unnamed import
	_ "github.com/go-sql-driver/mysql"
)

var dbConnPool *sql.DB

// StartConnectionPool func
func StartConnectionPool() error {
	fmt.Print("Initializing connection pool... ")

	err := OpenConnectionPool(
		viper.GetString("username") + ":" +
			viper.GetString("password") + "@" + viper.GetString("protocol") + "(" +
			viper.GetString("host") + ":" + viper.GetString("port") + ")/" +
			viper.GetString("database"))
	if err != nil {
		return err
	}

	color.Green("Done.\n")

	return nil
}

// OpenConnectionPool func
func OpenConnectionPool(dsn string) error {
	var err error

	if dbConnPool, err = sql.Open(viper.GetString("driver"), dsn); nil != err {
		return err
	}

	return dbConnPool.Ping()
}

// CloseConnectionPool func
func CloseConnectionPool() error {
	return dbConnPool.Close()
}
