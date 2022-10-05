package database

import (
	"database/sql"
	"log"
)

var Connector *sql.DB

func Connect(connectionString string) error {
	var err error
	Connector, err = sql.Open("mysql", connectionString)
	if err != nil {
		return err
	}

	log.Println("Connection was successful!!")
	return nil
}
