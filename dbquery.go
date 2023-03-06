package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/mateors/msql"
)

// Function Get
// get user
func GetUser(userEmail, userPassword string) ([]map[string]interface{}, error) {
	qs := fmt.Sprintf("SELECT Email, Password FROM user WHERE Email='%s' AND Password='%s';", userEmail, userPassword)
	row, err := msql.GetAllRowsByQuery(qs, db_user)
	FetchError(err)
	return row, nil
}

// getemail
func GetEmail(email string) ([]map[string]interface{}, error) {
	qs := fmt.Sprintf("SELECT Email FROM user WHERE Email='%s' AND status=1;", email)
	row, err := msql.GetAllRowsByQuery(qs, db_user)
	FetchError(err)
	return row, nil
}

// get signup user
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

// // get update password
func UpdatePassword(upPass, useremail string) (bool, error) {
	qs := fmt.Sprintf("UPDATE user SET Password = '%s' WHERE Email='%s'", upPass, useremail)
	row := msql.RawSQL(qs, db_user)
	return row, nil
}
