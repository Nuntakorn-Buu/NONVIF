package main

import (
	"database/sql"
	"fmt"
	"log"
)

var (
	db  *sql.DB
	err error
)

func database_user() {
	// Connect to database
	db, err = sql.Open("sqlite3", "./loginsystem.db")
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxOpenConns(1)
	log.Println("| Database User | Connection : Successful")
	fmt.Println(" ")
}
