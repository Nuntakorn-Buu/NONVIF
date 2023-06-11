// `package main` is the declaration of the package name.
// In Go, the `main` package is a special package that is used to build executable programs.
// It must contain a `main` function as the entry point of the program.
package main

// The `import` statement is used to import packages that are necessary for the program to run.
// In this case, the program is importing the following packages:
import (
	"fmt"     // Tools for manipulating the screen display. and reading data from the user You can use the Printf() function.
	"log"     // Use the message log (log) and don't forget the message. (error reporting)
	"net/url" // Used to manipulate and parse URLs, which are a form of characters used to address network resources such as website URLs.

	"github.com/mateors/msql" // It is a library for working with MySQL, SQLite databases in Go language.
)

// Getuser
// The function retrieves user data from a database based on their email and password.
func GetUser(userEmail, userPassword string) ([]map[string]interface{}, error) {
	qs := fmt.Sprintf("SELECT Email, Password FROM user WHERE Email='%s' AND Password='%s';", userEmail, userPassword)
	row, err := msql.GetAllRowsByQuery(qs, db_user)
	FetchError(err)
	return row, nil
}

// Getemail
// The function retrieves a list of email addresses from a database table for a given email address if it exists and has a status of 1.
func GetEmail(email string) ([]map[string]interface{}, error) {
	qs := fmt.Sprintf("SELECT Email FROM user WHERE Email='%s' AND status=1;", email)
	row, err := msql.GetAllRowsByQuery(qs, db_user)
	FetchError(err)
	return row, nil
}

// GetsignupUser
// The function inserts user data into a SQLite database table and returns the ID of the inserted record.
func SignupUser(name, email, username, mobile, password string) (int64, error) {
	data := make(url.Values)
	data.Set("table", "user")
	data.Set("dbtype", "sqlite3")
	data.Set("Name", name)
	data.Set("Email", email)
	data.Set("UserName", username)
	data.Set("Mobile", mobile)
	data.Set("Password", password)
	id, err := msql.InsertIntoAnyTable(data, db_user)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	fmt.Println("Successfully Inserted", id)
	return id, nil
}

// GetupdatePassword
// The function updates a user's password in the database.
func UpdatePassword(upPass, useremail string) (bool, error) {
	qs := fmt.Sprintf("UPDATE user SET Password = '%s' WHERE Email='%s'", upPass, useremail)
	row := msql.RawSQL(qs, db_user)
	return row, nil
}
