package main

//ทำการ import library ที่จำเป็น
import (
	"fmt"
	"net/http"
	"project-nonvif/view"
	"strconv"

	"github.com/gorilla/sessions"
)

// ประกาศตัวแปรต่าง ๆ
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

// สร้าง API ต่างๆ
// API Login
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
		err := loginView.Template.Execute(w, "Please give me right username or password")
		FetchError(err)
	}
}

// API forgot Password
func forgotPass(w http.ResponseWriter, _ *http.Request) {
	err := forgotPassView.Template.Execute(w, nil)
	FetchError(err)
}

// API forgot Password Authrization
var isEmail string

func forgotPassAuth(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.FormValue("forgotEmail")
	userEmail, err := GetEmail(email)
	FetchError(err)
	if len(userEmail) > 0 {
		isEmail = userEmail[0]["Email"].(string)
		if email == isEmail {
			err := fotgotAuthView.Template.Execute(w, "Check Your Email")
			emailSend(isEmail)
			FetchError(err)
		}
	} else {
		err := fotgotAuthErrorView.Template.Execute(w, "Your account does not exist")
		FetchError(err)
	}

}

// API forgot Code Verify
var randomNUM int = int(sixDigits())

func forgotCodeVerify(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(randomNUM)
	codeSt := r.FormValue("forgotEmail")
	codeint, err := strconv.ParseInt(codeSt, 10, 64)
	FetchError(err)
	if randomNUM == int(codeint) {
		err := updatePassView.Template.Execute(w, nil)
		FetchError(err)
	} else {
		fmt.Println("No")
	}

}

// API Check Password
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
func cameras(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "login-session")
	FetchError(err)
	_, ok := session.Values["username"]
	if ok {
		err := camerasView.Template.Execute(w, nil)
		FetchError(err)
	} else {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

}

// API Dashboard page
func dashboard(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "login-session")
	FetchError(err)
	_, ok := session.Values["username"]
	if ok {
		err := dashboardView.Template.Execute(w, nil)
		FetchError(err)
	} else {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

}

// API About page
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
func logout(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "login-session")
	FetchError(err)
	delete(session.Values, "username")
	session.Save(r, w)
	http.Redirect(w, r, "/login", http.StatusFound)
}

// API not Fount page
func notFount(w http.ResponseWriter, _ *http.Request) {
	err := notFountView.Template.Execute(w, nil)
	FetchError(err)
}

// Note สิ่งที่แก้

// API not Fount page
// func notFount(w http.ResponseWriter, r *http.Request) {
// 	err := notFountView.Template.Execute(w, nil)
// 	FetchError(err)
// }

// API forgot Password
// func forgotPass(w http.ResponseWriter, r *http.Request) {
// 	err := forgotPassView.Template.Execute(w, nil)
// 	FetchError(err)
// }
