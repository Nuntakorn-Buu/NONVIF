<p align="center">
	<img height="300" src="https://github.com/Nuntakorn-Buu/NONVIF/blob/main/view/front-end/asset/images/logo-3.png?raw=true">
</p>

# NONVIF

- การประยุกต์ใช้โปรโตคอล ONVIF เพื่อรองรับการจัดการกล้องไอพีและการสตรีมวิดีโอผ่านเว็บไซต์ .
- Application of the ONVIF Protocol to support IP camera management and video streaming through websites .

# NONVIF System

## ONVIF Profile T

- Streaming Video
- Metadata

---

## Features Projects

- Home page
- Cameras page
- Dashboard page
- About page
- Portfolio page
- Add camera
- Remove camera
- Signup
- Login
- Logout
- Forget Password

---

# Details

## Recommend

- Ubuntu 20.04.5 LTS
  - Installation method ---> [Link](https://youtu.be/QKn5U2esuRk)
- Golang | version 1.19
  - Installation method ---> [Link](https://www.dropbox.com/scl/fi/9clhqzmcuk0gcjgby2h38/Golang.paper?dl=0&rlkey=pipvjgui4256g0wn4sp2aoklj)

## Dev Tool

- VScode or Visual Studio Code version | 1.76
  - Installation method ---> [Link](https://code.visualstudio.com/)
- DB Browser for SQLite | version 3.12.2
  - Installation method ---> [Link](https://sqlitebrowser.org/dl/)
    #### or
  - Github sqlitebrowser ---> [Link](https://github.com/sqlitebrowser/sqlitebrowser)

## Hardware Accessories

- At least 2 IP cameras
  #### or
- 1 IP camera and Programe IP Webcam | version 1.16.6.783
  - Installation method ---> [Link](https://play.google.com/store/apps/details?id=com.pas.webcam&hl=th)

---

## Backend

- HTTP Middleware ( [gorilla/mux](https://github.com/gorilla/mux) )
- Login Session ( [gorilla/sessions](https://github.com/gorilla/sessions) )
- Database Sqlite ( [go-sqlite3](https://github.com/mattn/go-sqlite3) & [msql](https://github.com/mateors/msql) )

## Frontend

- Server side templating ( [Go HTML Templates](https://pkg.go.dev/html/template) )
- Frontend `Pure ( HTML , CSS )`

## Framework

- w3.css
  - Installation method ---> [Link](https://www.w3schools.com/w3css/)
- Font-Awesome
  - Home page ---> [Link](https://fontawesome.com/)
  - Installation method ---> [Link](https://cdnjs.com/libraries/font-awesome)

## Modules

```go
require (
    github.com/gorilla/mux v1.8.0
	github.com/gorilla/sessions v1.2.1
	github.com/mateors/msql v0.0.0-20221130043645-280860a386a7
	github.com/mattn/go-sqlite3 v1.14.16
)

// indirect
require (
    github.com/gorilla/securecookie v1.1.1
)
```

---

## Step Install Repositories !!!

### Step 1 :

- ติดตั้งตามที่ได้ [Recommend](#recommend) ไว้ก่อนตามที่ระบุไว้ด้านบน

### Step 2 :

- ทำการ clone Repositories นี้ จาก [Nuntakorn-Buu/NONVIF](https://github.com/Nuntakorn-Buu/NONVIF)

  หรือ ใช้คำสั่ง

  ```bash
  git clone https://github.com/Nuntakorn-Buu/NONVIF.git
  ```

### Step 3 :

- หลังจากทำการ clone Repositories มาแล้ว. ให้ใช้คำสั่ง `< go get >` เพื่อติตดั้ง [Modules](#modules) ที่จำเป็นต่าง ๆ
  ```go
  go get
  ```

### Step 4 :

- หลังจากทำตาม [Step 3 :](#step-3) เสร็จสิ้น. ต่อมาจะเป็นการใช้คำสั่ง `< go run. >` เพื่อทำการ Run Project นี้
  ```go
  go run .
  ```

---
