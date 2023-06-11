// `package main` is the declaration of the package name.
// In Go, the `main` package is a special package that is used to build executable programs.
// It must contain a `main` function as the entry point of the program.
package main

// The `import` statement is used to import packages that are necessary for the program to run.
// In this case, the program is importing the following packages:
import (
	"fmt"                 // Tools for manipulating the screen display. and reading data from the user You can use the Printf() function.
	"net/http"            // Used to create a Go web server, with functions and data structures that provide services for handling network connections and HTTP traffic, such as the HandleFunc() function.
	"project-nonvif/view" // It's a self-written package. contains data structures and functions related to web page rendering.

	"github.com/gorilla/mux"      // It helps to write server based web applications. It has functions and data structures that help manage web routing and passing data through parameters.
	"github.com/gorilla/sessions" // It helps to manage user sessions in the application. i.e. storing and reading values from user sessions.

	_ "github.com/mattn/go-sqlite3" // An SQLite database connection package in Go that can be used to work with SQLite databases.
)

// `var store = sessions.NewCookieStore([]byte("secret-password"))` is creating a new cookie store for session management.
// The `[]byte("secret-password")` parameter is used to encrypt and decrypt the session data stored in the cookie.
// This ensures that the session data is secure and cannot be tampered with by unauthorized users.
var store = sessions.NewCookieStore([]byte("secret-password"))

// The `init()` function calls two other functions related to database setup in Go.
func init() {
	database_user()    // run the database user
	database_cameras() // run the database cameras
}

// This function sets up a router and defines various routes and views for a web application.
func main() {
	// These lines of code are creating new views for different web pages of the application.
	// Each view is associated with an HTML file located in the "view/front-end/" directory.
	// These views will be used later in the code to render the corresponding web pages when the user navigates to them.
	homeView = view.NewView("view/front-end/index.html")
	aboutView = view.NewView("view/front-end/about.html")
	camerasView = view.NewView("view/front-end/cameras.html")
	dashboardView = view.NewView("view/front-end/dashboard.html")
	portfolioView = view.NewView("view/front-end/portfolio.html")
	notFountView = view.NewView("view/front-end/notfount.html")
	loginView = view.NewView("view/front-end/login.html")
	signupView = view.NewView("view/front-end/signin.html")
	forgotPassView = view.NewView("view/front-end/forgotpass.html")
	fotgotAuthView = view.NewView("view/front-end/forgetauth.html")
	fotgotAuthErrorView = view.NewView("view/front-end/forgotAuthError.html")
	updatePassView = view.NewView("view/front-end/updatepass.html")

	// `r := mux.NewRouter()` is creating a new router using the Gorilla mux package.
	// This router will be used to define various routes and views for a web application.
	// The mux package provides functions and data structures that help manage web routing and passing data through parameters.
	r := mux.NewRouter()
	r.PathPrefix("/asset/").Handler(http.StripPrefix("/asset/", http.FileServer(http.Dir("view/front-end/asset"))))
	r.NotFoundHandler = http.HandlerFunc(notFount)
	r.HandleFunc("/login", login)
	r.HandleFunc("/loginauth", loginAuth)
	r.HandleFunc("/", home)
	r.HandleFunc("/about", about)
	r.HandleFunc("/cameras", cameras)
	r.HandleFunc("/dashboard", dashboard)
	r.HandleFunc("/portfolio", portfolio)
	r.HandleFunc("/signup", signup)
	r.HandleFunc("/signupauth", signupAuth)
	r.HandleFunc("/logout", logout)
	r.HandleFunc("/forgot_pass", forgotPass)
	r.HandleFunc("/forgot_pass_auth", forgotPassAuth)
	r.HandleFunc("/code_verify", forgotCodeVerify)
	r.HandleFunc("/checkpass", checkPass)
	r.HandleFunc("/add_camera", addCamera).Methods("POST")
	r.HandleFunc("/remove_camera", removeCamera).Methods("POST")
	r.HandleFunc("/metadatas", metadatas).Methods("POST")

	// `fmt.Println("Listening port : 9000")` is printing the message "Listening port : 9000" to the console.
	// This message is printed when the program starts running and indicates that the program is listening on port 9000 for incoming requests.
	fmt.Println("Listening port : 9000")

	// `http.ListenAndServe(":9000", r)` is starting a web server that listens for incoming HTTP requests on port 9000 and uses the router `r` to handle those requests.
	// It is the final step in setting up a web application in Go.
	http.ListenAndServe(":9000", r)
}
