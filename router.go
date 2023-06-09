package main

//ทำการ import library ที่จำเป็น
import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"project-nonvif/view"
	"strconv"

	"github.com/inspii/onvif"

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

// สร้างตัวแปรเก็บข้อมูลกล้อง
var cameraURLs = []string{}

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
		// ดึงข้อมูลกล้องจากฐานข้อมูล cameras
		rows, err := db_cameras.Query("SELECT url FROM cameras")
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

		// เรียกใช้ฟังก์ชัน metadatas เพื่อแสดงข้อมูลเมื่อกดปุ่ม Show Data Cameras
		if r.Method == http.MethodPost {
			if r.FormValue("action") == "showData" {
				metadatas(w, r)
			}
		}

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
		// ดึงข้อมูลเซ็นเซอร์จากฐานข้อมูล cameras
		rows, err := db_cameras.Query("SELECT url FROM cameras")
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

// API Add Camera
func addCamera(w http.ResponseWriter, r *http.Request) {
	cameraURL := r.FormValue("cameraURL")

	// เพิ่มข้อมูลกล้องลงในฐานข้อมูล cameras
	insertStmt := `INSERT INTO cameras (url) VALUES (?)`
	_, err := db_cameras.Exec(insertStmt, cameraURL)
	if err != nil {
		log.Println(err)
		return
	}

	// ส่งตอบกลับว่าเพิ่มกล้องสำเร็จ
	w.WriteHeader(http.StatusOK)
}

// API Remove Camera
func removeCamera(w http.ResponseWriter, r *http.Request) {
	cameraURL := r.FormValue("cameraURL")

	// ลบข้อมูลกล้องออกจากฐานข้อมูล cameras
	deleteStmt := `DELETE FROM cameras WHERE url = ?`
	_, err := db_cameras.Exec(deleteStmt, cameraURL)
	if err != nil {
		log.Println(err)
		return
	}

	// ส่งตอบกลับว่าลบกล้องสำเร็จ
	w.WriteHeader(http.StatusOK)
}

// API metadatas
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

	// สร้างข้อมูล JSON จาก di,users,dateandtime,cap,hostname และ networkprotocols
	data := struct {
		DeviceInformation interface{}
		Users             interface{}
		SystemDateAndTime interface{}
		Capabilities      interface{}
		Hostname          interface{}
		NetworkProtocols  interface{}
	}{
		DeviceInformation: di,
		Users:             users,
		SystemDateAndTime: dateandtime,
		Capabilities:      cap,
		Hostname:          hostname,
		NetworkProtocols:  networkprotocols,
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

	// ส่งข้อมูลกลับไปยังหน้าเว็บในรูปแบบ JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
