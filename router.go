// `package main` is the declaration of the package name.
// In Go, the `main` package is a special package that is used to build executable programs.
// It must contain a `main` function as the entry point of the program.
package main

// The `import` statement is used to import packages that are necessary for the program to run.
// In this case, the program is importing the following packages:
import (
	"database/sql"        // Used to work with databases. It supports connection and data management of multiple databases such as MySQL, PostgreSQL, SQLite, etc.
	"encoding/json"       // Used to encode and decode data in JSON (JavaScript Object Notation) format. It helps to convert Go data to JSON and convert JSON back to Go data.
	"fmt"                 // Tools for manipulating the screen display. and reading data from the user You can use the Printf() function.
	"log"                 // Use the message log (log) and don't forget the message. (error reporting)
	"net/http"            // Used to create a Go web server, with functions and data structures that provide services for handling network connections and HTTP traffic, such as the HandleFunc() function.
	"project-nonvif/view" // It's a self-written package. contains data structures and functions related to web page rendering.
	"strconv"             // Used to convert numeric and string values. This package provides functions that facilitate data conversions related to numeric manipulation.

	"github.com/inspii/onvif" // It is a library used to connect and control ONVIF (Open Network Video Interface Forum) standard peripherals. It contains functions and data structures that help connect and communicate with ONVIF devices.

	"github.com/gorilla/sessions" //It helps to manage user sessions in the application. i.e. storing and reading values from user sessions.
)

// The above code is declaring multiple variables of type *view.View in Go language.
// These variables are likely to be used to store different views or templates for a web application.
// The names of the variables suggest that they may correspond to different pages of the web application
// such as home, about, cameras, dashboard, portfolio, login, signup, forgot password, etc.
var (
	homeView            *view.View
	aboutView           *view.View
	camerasView         *view.View
	dashboardView       *view.View
	portfolioView       *view.View
	notFountView        *view.View
	loginView           *view.View
	signupView          *view.View
	forgotPassView      *view.View
	fotgotAuthView      *view.View
	fotgotAuthErrorView *view.View
	updatePassView      *view.View
)

// The above code is declaring an empty slice of strings named `cameraURLs` in the Go programming language.
var cameraURLs = []string{}

// Build various APIs

// API Login
// The login function checks if a user is already logged in and redirects them if they are, otherwise it displays the login view.
func login(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "login-session")
	FetchError(err)
	_, ok := session.Values["username"]
	if ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if !ok {
		err := loginView.Template.Execute(w, nil)
		FetchError(err)
	}
}

// API Login Authorization
// This function handles user authentication and session creation for a login page in a Go web application.
func loginAuth(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.FormValue("email")
	password := r.FormValue("password")
	user, err := GetUser(email, password)
	FetchError(err)
	if len(user) > 0 {
		session, err := store.Get(r, "login-session")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		session.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   86400 * 30,
			HttpOnly: true,
		}
		session.Values["username"] = "username"
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		err := loginView.Template.Execute(w, "Please give me right email or password")
		FetchError(err)
	}
}

// API forgot Password
// The function serves the forgot password view template to the user.
func forgotPass(w http.ResponseWriter, _ *http.Request) {
	err := forgotPassView.Template.Execute(w, nil)
	FetchError(err)
}

// API forgot Password Authrization
// The function handles the authentication process for a forgotten password request,
// checking if the email exists and sending a code to the user's email if it does.
var isEmail string

func forgotPassAuth(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.FormValue("forgotEmail")
	userEmail, err := GetEmail(email)
	FetchError(err)
	if len(userEmail) > 0 {
		isEmail = userEmail[0]["Email"].(string)
		if email == isEmail {
			type Data struct {
				Code string
			}
			data := Data{
				Code: strconv.Itoa(randomNUM),
			}
			err := fotgotAuthView.Template.Execute(w, data)
			emailSend(isEmail)
			FetchError(err)
		}
	} else {
		err := fotgotAuthErrorView.Template.Execute(w, "Your account does not exist")
		FetchError(err)
	}
}

// API forgot Code Verify
// This function verifies a code entered by the user and redirects them to a password update page if the code is correct,
// otherwise it redirects them back to the forgot password page.
var randomNUM int = int(sixDigits())

func forgotCodeVerify(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(randomNUM)
	codeSt := r.FormValue("forgotEmail")
	codeint, err := strconv.ParseInt(codeSt, 10, 64)
	FetchError(err)
	if randomNUM == int(codeint) {
		data := struct {
			Code int
		}{
			Code: randomNUM,
		}
		err := updatePassView.Template.Execute(w, data)
		FetchError(err)
	} else {
		http.Redirect(w, r, "/forgot_pass", http.StatusSeeOther)
		fmt.Println("-------------------- | No")
	}
}

// API Check Password
// The function checks if two password inputs match and updates the password if they do, otherwise it displays an error message.
func checkPass(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	pass1 := r.FormValue("pass1")
	pass2 := r.FormValue("pass2")
	if pass1 == pass2 {
		UpdatePassword(pass1, isEmail)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {
		updatePassView.Template.Execute(w, "Please Make sure Your password Both of same")
	}
}

// API Home page
// The function checks if a user is logged in and redirects them to the login page if not.
func home(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "login-session")
	FetchError(err)
	_, ok := session.Values["username"]
	if ok {
		err := homeView.Template.Execute(w, nil)
		FetchError(err)
	} else {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

// API Cameras page
// This function retrieves camera information from the database and displays it on a web page.
// If the user is not logged in to the system will be redirected to the login page.
func cameras(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "login-session")
	FetchError(err)
	_, ok := session.Values["username"]
	if ok {
		rows, err := db_cameras.Query("SELECT url FROM cameras") // Retrieve camera data from cameras database.
		if err != nil {
			log.Println(err)
			return
		}
		defer rows.Close()

		var cameraURLs []string
		for rows.Next() {
			var url sql.NullString
			err = rows.Scan(&url)
			if err != nil {
				log.Println(err)
				return
			}
			if url.Valid {
				cameraURLs = append(cameraURLs, url.String)
			} else {
				cameraURLs = append(cameraURLs, "")
			}
		}
		if err = rows.Err(); err != nil {
			log.Println(err)
			return
		}

		err = camerasView.Template.Execute(w, struct {
			CameraURLs []string
		}{
			CameraURLs: cameraURLs,
		})
		FetchError(err)

	} else {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

// API Dashboard page
// The function retrieves camera data from a database and displays it on a dashboard page if the user is logged in,
// otherwise it redirects to the login page.
func dashboard(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "login-session")
	FetchError(err)
	_, ok := session.Values["username"]
	if ok {
		rows, err := db_cameras.Query("SELECT url FROM cameras") // Retrieve camera data from cameras database.
		if err != nil {
			log.Println(err)
			return
		}
		defer rows.Close()

		var cameraURLs []string
		for rows.Next() {
			var url sql.NullString
			err := rows.Scan(&url)
			if err != nil {
				log.Println(err)
				return
			}
			if url.Valid {
				cameraURLs = append(cameraURLs, url.String)
			} else {
				cameraURLs = append(cameraURLs, "")
			}
		}
		if err := rows.Err(); err != nil {
			log.Println(err)
			return
		}

		err = dashboardView.Template.Execute(w, struct {
			CameraURLs []string
		}{
			CameraURLs: cameraURLs,
		})
		FetchError(err)
	} else {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

// API About page
// This function checks if a user is logged in and redirects them to the login page if they are not.
func about(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "login-session")
	FetchError(err)
	_, ok := session.Values["username"]
	if ok {
		err := aboutView.Template.Execute(w, nil)
		FetchError(err)
	} else {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

// API Portfolio page
// This function checks if a user is logged in and redirects them to the login page if not,
// otherwise it displays the portfolio view.
func portfolio(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "login-session")
	FetchError(err)
	_, ok := session.Values["username"]
	if ok {
		err := portfolioView.Template.Execute(w, nil)
		FetchError(err)
	} else {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

// API Signup page
// The function checks if a user is logged in and redirects them to the homepage if they are,
// otherwise it displays the signup view.
func signup(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "login-session")
	FetchError(err)
	_, ok := session.Values["username"]
	if !ok {
		err := signupView.Template.Execute(w, nil)
		FetchError(err)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

}

// API Signup Authorization page
// This function handles the signup authentication process by parsing form data, creating a new user,
// and setting a session cookie before redirecting to the homepage.
func signupAuth(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.FormValue("fullName")
	email := r.FormValue("email")
	userName := r.FormValue("userName")
	mobile := r.FormValue("mobileNumber")
	password := r.FormValue("Password")

	SignupUser(name, email, userName, mobile, password)
	session, err := store.Get(r, "login-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 30,
		HttpOnly: true,
	}
	session.Values["username"] = "username"
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

// API Logout page
// The function logs out a user by deleting their session and redirecting them to the login page.
func logout(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "login-session")
	FetchError(err)
	delete(session.Values, "username")
	session.Save(r, w)
	http.Redirect(w, r, "/login", http.StatusFound)
}

// API not Fount page
// This function handles a "not found" error by executing a template and writing it to the response.
func notFount(w http.ResponseWriter, _ *http.Request) {
	err := notFountView.Template.Execute(w, nil)
	FetchError(err)
}

// API Add Camera
// The function adds a camera URL to a database table and sends a success response.
func addCamera(w http.ResponseWriter, r *http.Request) {
	cameraURL := r.FormValue("cameraURL")
	insertStmt := `INSERT INTO cameras (url) VALUES (?)` // Add camera information to the cameras database.
	_, err := db_cameras.Exec(insertStmt, cameraURL)
	if err != nil {
		log.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK) // sends a reply that the camera was added successfully
}

// API Remove Camera
// This function removes camera data from a database based on the camera URL provided in an HTTP request.
func removeCamera(w http.ResponseWriter, r *http.Request) {
	cameraURL := r.FormValue("cameraURL")
	deleteStmt := `DELETE FROM cameras WHERE url = ?` // delete camera data from cameras database
	_, err := db_cameras.Exec(deleteStmt, cameraURL)
	if err != nil {
		log.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK) // sends a reply that the camera was added successfully
}

// API metadatas
// The function retrieves metadata information from an ONVIF camera and sends it back to the web page in JSON response format.
func metadatas(w http.ResponseWriter, r *http.Request) {
	cameraURL := r.FormValue("cameraURL")
	camera := onvif.NewOnvifDevice(cameraURL)
	camera.Auth("admin", "admin") // "user", "password"

	err := camera.Initialize()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	di, err := camera.Device.GetDeviceInformation()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	users, err := camera.Device.GetUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dateandtime, err := camera.Device.GetSystemDateAndTime()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cap, err := camera.Device.GetCapabilities()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	hostname, err := camera.Device.GetHostname()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	networkprotocols, err := camera.Device.GetNetworkProtocols()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	discovery, err := camera.Device.GetDiscoveryMode()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	servicecap, err := camera.Device.GetServiceCapabilities()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Generate JSON data from di,users,dateandtime,cap,hostname networkprotocols, discovery and servicecap.
	data := struct {
		DeviceInformation   interface{}
		Users               interface{}
		SystemDateAndTime   interface{}
		Capabilities        interface{}
		Hostname            interface{}
		NetworkProtocols    interface{}
		DiscoveryMode       interface{}
		ServiceCapabilities interface{}
	}{
		DeviceInformation:   di,
		Users:               users,
		SystemDateAndTime:   dateandtime,
		Capabilities:        cap,
		Hostname:            hostname,
		NetworkProtocols:    networkprotocols,
		DiscoveryMode:       discovery,
		ServiceCapabilities: servicecap,
	}
	fmt.Println("Data From GetDeviceInformation :", di)
	fmt.Println("")
	fmt.Println("Data From GetUsers :", users)
	fmt.Println("")
	fmt.Println("Data From GetSystemDateAndTime :", dateandtime)
	fmt.Println("")
	fmt.Println("Data From GetCapabilities :", cap)
	fmt.Println("")
	fmt.Println("Data From GetHostname :", hostname)
	fmt.Println("")
	fmt.Println("Data From GetNetworkProtocols :", networkprotocols)
	fmt.Println("")
	fmt.Println("Data From GetDiscoveryMode :", discovery)
	fmt.Println("")
	fmt.Println("Data From GetServiceCapabilities :", servicecap)
	fmt.Println("")

	// Sends data back to the web page in JSON response format.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
