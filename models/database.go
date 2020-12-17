package models

import (
	"database/sql"
	"fmt"
	"os"
)

type Database struct {
	host       string
	port       string
	name       string
	username   string
	password   string
	connection *sql.DB
}

func NewDatabase() *Database {
	return &Database{
		host:     os.Getenv("DATABASE_HOST"),
		port:     os.Getenv("DATABASE_PORT"),
		name:     os.Getenv("DATABASE_NAME"),
		username: os.Getenv("DATABASE_USERNAME"),
		password: os.Getenv("DATABASE_PASSWORD"),
	}
}

func (db *Database) Connect() {
	var err error
	db.connection, err = sql.Open("mysql", db.username+":"+db.password+"@tcp("+db.host+":"+db.port+")/"+db.name)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Connection to database successful")
	defer db.connection.Close()
}
