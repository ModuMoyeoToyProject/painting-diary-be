package db

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

// Use a mysql.Config object to build the DSN. Please note that
// it's not recommended to store the password directly in the
// code for production usage.
func withConfig() {
	config := mysql.NewConfig()
	config.User = "0a0aavo8l9ku8hkzr3r3"
	config.Passwd = "pscale_pw_qSMSo9izUx0s0ljYPac1BF7zGEUvVHcgBqve8j9y9RI"
	config.Net = "tcp"
	config.Addr = "aws.connect.psdb.cloud"
	config.DBName = "carrot_auction"
	config.TLSConfig = "true"

	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
}

func Connection() {
	withConfig()
	fmt.Println("Successfully connected to PlanetScale with configuration object!")
}