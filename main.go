package main

// ทำการ import library ที่จำเป็นต้องใช้มา
import (
	"fmt"
	"net/http"
	"project-nonvif/view"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"

	_ "github.com/mattn/go-sqlite3"
)

// สร้างตัวแปรที่จะเก็บข้อมูล
var store = sessions.NewCookieStore([]byte("secret-password"))

// สร้างฟังก์ชั่นสำหรับเรียกใช้ database
func init() {
	database()
}

func main() {
	homeView = view.NewView("view/front-end/index.html")
	aboutView = view.NewView("view/front-end/about.html")
	portfolioView = view.NewView("view/front-end/portfolio.html")
	notFountView = view.NewView("view/front-end/notfount.html")
	loginView = view.NewView("view/front-end/login.html")
	signupView = view.NewView("view/front-end/signin.html")
	forgotPassView = view.NewView("view/front-end/forgotpass.html")
	fotgotAuthView = view.NewView("view/front-end/forgetauth.html")
	fotgotAuthErrorView = view.NewView("view/front-end/forgotAuthError.html")
	updatePassView = view.NewView("view/front-end/updatepass.html")

	r := mux.NewRouter()
	r.PathPrefix("/asset/").Handler(http.StripPrefix("/asset/", http.FileServer(http.Dir("view/front-end/asset"))))
	r.NotFoundHandler = http.HandlerFunc(notFount)
	r.HandleFunc("/login", login)
	r.HandleFunc("/loginauth", loginAuth)
	r.HandleFunc("/", home)
	r.HandleFunc("/about", about)
	r.HandleFunc("/portfolio", portfolio)
	r.HandleFunc("/signup", signup)
	r.HandleFunc("/signupauth", signupAuth)
	r.HandleFunc("/logout", logout)
	r.HandleFunc("/forgot_pass", forgotPass)
	r.HandleFunc("/forgot_pass_auth", forgotPassAuth)
	r.HandleFunc("/code_verify", forgotCodeVerify)
	r.HandleFunc("/checkpass", checkPass)

	fmt.Println("Listening port :9000")

	http.ListenAndServe(":9000", r)
}