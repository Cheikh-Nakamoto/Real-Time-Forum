package database

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	db *sql.DB
}

func (d *Database) GetDatabase() *sql.DB {
	return d.db
}

func (d *Database) Close() {
	err := d.db.Close()
	if err != nil {
		fmt.Println("Error closing database connection:", err)
	}
}

func NewDatabase() (*Database, error) {
	db, err := sql.Open("sqlite3", "./internal/metier/database/mytemplate.db")
	if err != nil {
		return nil, err
	}
	err = db.Ping() // Check connection with database
	if err != nil {
		err := db.Close() // Ensure the database is closed if ping fails
		if err != nil {
			return nil, err
		}
		return nil, err
	}
	fmt.Println("Connected to database")
	return &Database{db: db}, nil
}

func InitDatabase() (*Database, error) {
	db, err := NewDatabase()
	file, err := ioutil.ReadFile("./internal/metier/database/migrations/tables.up.sql")
	if err != nil {
		fmt.Println("Could not read SQL file: ", err)
		return nil, err
	}

	cmds := strings.Split(string(file), ";")
	for _, cmd := range cmds {
		cmd = strings.TrimSpace(cmd)
		if cmd == "" {
			continue
		}
		_, err = db.db.Exec(cmd)
		if err != nil {
			fmt.Println("Error executing SQL command:", cmd)
			return nil, err
		}
	}
	return db, nil
}
