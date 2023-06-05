package main

import (
	"database/sql"
	"fmt"
	"log"
)

var (
	db_user    *sql.DB
	db_cameras *sql.DB
	err        error
)

func database_user() {
	// Connect to user database
	db_user, err = sql.Open("sqlite3", "./loginsystem.db")
	if err != nil {
		log.Fatal(err)
	}
	db_user.SetMaxOpenConns(1)
	log.Println("| Database User | Connection : Successful")
}

func database_cameras() {
	// Connect to cameras database
	db_cameras, err = sql.Open("sqlite3", "./cameras.db")
	if err != nil {
		log.Fatal(err)
	}
	// Create cameras table if not exists
	createTableStmt := `
		CREATE TABLE IF NOT EXISTS cameras (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			url TEXT
		)
	`
	_, err = db_cameras.Exec(createTableStmt)
	if err != nil {
		log.Fatal(err)
	}

	db_cameras.SetMaxOpenConns(1)
	log.Println("| Database Cameras | Connection : Successful")
	fmt.Println(" ")
}
