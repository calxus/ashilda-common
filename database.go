package models

import (
	"database/sql"
	"fmt"
	"os"
)

type Database struct {
	Host       string
	Port       string
	Name       string
	Username   string
	Password   string
	Connection *sql.DB
}

func NewDatabase() *Database {
	return &Database{
		Host:     os.Getenv("DATABASE_HOST"),
		Port:     os.Getenv("DATABASE_PORT"),
		Name:     os.Getenv("DATABASE_NAME"),
		Username: os.Getenv("DATABASE_USERNAME"),
		Password: os.Getenv("DATABASE_PASSWORD"),
	}
}

func (db *Database) Connect() {
	var err error
	db.Connection, err = sql.Open("mysql", db.Username+":"+db.Password+"@tcp("+db.Host+":"+db.Port+")/"+db.Name)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Connection to database successful")
	defer db.Connection.Close()
}
