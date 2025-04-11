package config

import (
	"database/sql"
	// Here we need to import the driver anonymously
	_ "github.com/go-sql-driver/mysql"
)

const (
	driverName = "mysql"
	connectionString = "root:aditya@tcp(127.0.0.1:3306)/adityadb"
)

func CreateConnection() (*sql.DB, error) {
	db, err := sql.Open(driverName, connectionString)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}