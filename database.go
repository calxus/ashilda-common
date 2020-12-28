package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // Driver for MySQL database connection
)

// Database type holds the details for the connection and the connection itself
type Database struct {
	Host       string
	Port       string
	Name       string
	Username   string
	Password   string
	Connection *sql.DB
}

// NewDatabase method to construct the database type
func NewDatabase() *Database {
	return &Database{
		Host:     os.Getenv("DATABASE_HOST"),
		Port:     os.Getenv("DATABASE_PORT"),
		Name:     os.Getenv("DATABASE_NAME"),
		Username: os.Getenv("DATABASE_USERNAME"),
		Password: os.Getenv("DATABASE_PASSWORD"),
	}
}

// Connect method attempts to establish a connection to the database
func (db *Database) Connect() {
	var err error
	db.Connection, err = sql.Open("mysql", db.Username+":"+db.Password+"@tcp("+db.Host+":"+db.Port+")/"+db.Name)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Connection to database successful")
}

// ExecuteSelect method executes a select statement and returns the rows
func (db *Database) ExecuteSelect(ds *DatabaseSelect) (*sql.Rows, error) {
	rows, err := db.Connection.Query(ds.Generate())
	if err != nil {
		log.Fatal(err.Error())
	}
	return rows, err
}

// ExecuteUpdate method executes an SQL update statement
func (db *Database) ExecuteUpdate(du *DatabaseUpdate) error {
	_, err := db.Connection.Exec(du.Generate())
	if err != nil {
		log.Fatal(err.Error())
	}
	return err
}

// ExecuteDelete method executes an SQL delete statement
func (db *Database) ExecuteDelete(dd *DatabaseDelete) error {
	_, err := db.Connection.Exec(dd.Generate())
	if err != nil {
		log.Fatal(err.Error())
	}
	return err
}

// ExecuteInsert method executes an SQL insert statement
func (db *Database) ExecuteInsert(di *DatabaseInsert) error {
	_, err := db.Connection.Exec(di.Generate())
	if err != nil {
		log.Fatal(err.Error())
	}
	return err
}
