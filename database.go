// `package main` is the declaration of the package name.
// In Go, the `main` package is a special package that is used to build executable programs.
// It must contain a `main` function as the entry point of the program.
package main

// The `import` statement is used to import packages that are necessary for the program to run.
// In this case, the program is importing the following packages:
import (
	"database/sql" // Used to work with databases. It supports connection and data management of multiple databases such as MySQL, PostgreSQL, SQLite, etc.
	"fmt"          // Tools for manipulating the screen display. and reading data from the user You can use the Printf() function.
	"log"          // Use the message log (log) and don't forget the message. (error reporting)
)

// This is declaring three variables `db_user`, `db_cameras`, and `err`.
// `db_user` and `db_cameras` are pointers to a `sql.DB` struct, which is used to connect to a database.
// `err` is a variable of type `error`, which is used to store any errors that occur during the execution of the program.
var (
	db_user    *sql.DB
	db_cameras *sql.DB
	err        error
)

// This function connects to a user database in SQLite3 and sets the maximum number of open connections to 1.
func database_user() {
	// Connect to user database
	db_user, err = sql.Open("sqlite3", "./loginsystem.db")
	if err != nil {
		log.Fatal(err)
	}
	db_user.SetMaxOpenConns(1)
	log.Println("| Database User | Connection : Successful")
}

// The function connects to a SQLite database for cameras and creates a table if it does not exist.
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
