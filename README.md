![Logo](https://github.com/Nuntakorn-Buu/NONVIF/blob/main/view/front-end/asset/images/logo-5.png?raw=true)

# 🗂 NONVIF Project.

[TH] : การประยุกต์ใช้โปรโตคอล ONVIF เพื่อรองรับการจัดการกล้องไอพีและการสตรีมวิดีโอผ่านเว็บไซต์ .

[ENG] : Application of the ONVIF Protocol to support IP camera management and video streaming through websites .

### Project Description !!!

This project is an implementation of the ONVIF protocol on a web interface. It involves creating a web page to support video streaming from an Android application called 'IP Webcam.' This app simulates an IP camera that supports the ONVIF protocol and uses ONVIF protocol functions to retrieve various camera data from IP Webcam and display it on the web page

โครงการนี้เป็นการประยุกต์ใช้โปรโตคอล ONVIF ที่หน้าเว็บไซต์. โดยเป็นการสร้างหน้าเว็บขึ้นมาเพื่อใช้สำหรับรองรับการสตรีมภาพวิดีโอจากแอพพลิเคชันบนแอนดรอย์ ชื่อ "IP Webcam" ซึ่งเป็นแอพที่จำลองกล้องไอพีที่รองรับONVIF Protocol และใช้ Func ของ ONVIF Protocol เพื่อดึงข้อมูลต่าง ๆ กล้องจาก IP Webcam มาแสดงออกที่หน้าเว็บไซต์

---

## 🛠 Skills

The essential coding skills required for this project.

- 👉🏻 | Golang
- 👉🏻 | SQL
- 👉🏻 | HTML
- 👉🏻 | CSS
- 👉🏻 | Javascript

---

## 🔗 Links Documentations

| Docs             | Link                                                                                                                              |
| ---------------- | --------------------------------------------------------------------------------------------------------------------------------- |
| ➡️ Dropbox       | ➡️ [LINK](https://www.dropbox.com/sh/jysm351v3pfmnyr/AABIzDHCTdY0mFt5ScdX-Uzba?dl=0)                                              |
| ➡️ Task List     | ➡️ [LINK](https://nuntakorn-sp.notion.site/1d183abc6dea4b3ba570bf8726195ae0?v=c5632d2adb3d40deaa79cb392fc82038&pvs=4)             |
| ➡️ Work Plan 1   | ➡️ [LINK](https://www.dropbox.com/scl/fi/c98cdfjwiek2gxylp7h1y/Technical-Report-Non.paper?dl=0&rlkey=99122h15yh6enbrmwwgfkf79g)   |
| ➡️ Gantt Chart 1 | ➡️ [LINK](https://docs.google.com/spreadsheets/d/1imVBrtQ3R5AznIkBnEeeSc24qYHkcrq0weeTZ88HP44/edit#gid=611801224)                 |
| ➡️ Work Plan 2   | ➡️ [LINK](https://www.dropbox.com/scl/fi/izdrnq87zf7v8mjhbfntg/Work-Plan-Work-Log-Non.paper?dl=0&rlkey=dqaetj5gil8auaxsuqumg0p7s) |
| ➡️ Gantt Chart 2 | ➡️ [LINK](https://docs.google.com/spreadsheets/d/1imVBrtQ3R5AznIkBnEeeSc24qYHkcrq0weeTZ88HP44/edit#gid=902234290)                 |

---

## 📢 Features Projects

| Pages             | Systemes                |
| ----------------- | ----------------------- |
| ➡️ Home page      | ➡️ Sign up              |
| ➡️ Cameras page   | ➡️ Login                |
| ➡️ Dashboard page | ➡️ Logout               |
| ➡️ About page     | ➡️ Forget Password      |
| ➡️ Portfolio page | ➡️ Add or Remove camera |

### The ONVIF services utilized in this project.

| Function                   | Description                                                              |
| -------------------------- | ------------------------------------------------------------------------ |
| ✅GetDeviceInformation()   | This operation gets basic device information from the device.            |
| ✅GetUsers()               | This operation lists registered users and their credentials on a device. |
| ✅GetSystemDateAndTime()   | This operation gets the device system date and time.                     |
| ✅GetCapabilities()        | This operation will get various values ​​about Capabilities              |
| ✅GetHostname()            | This operation is used by an endpoint to get the hostname from a device. |
| ✅GetNetworkProtocols()    | This operation gets defined network protocols from a device.             |
| ✅GetDiscoveryMode()       | This operation gets the discovery mode of a device.                      |
| ✅GetServiceCapabilities() | This operation returns the capabilities of the PTZ service.              |

---

## 📷 Demo

Clip Video : Please click on the camera image to view the Demo.

[![Test Project](https://github.com/Nuntakorn-Buu/NONVIF/blob/main/view/front-end/asset/images/camera-1.png?raw=true)](https://youtu.be/6YH1yIA-MCc)

<iframe width="400" height="0" src="https://youtu.be/6YH1yIA-MCc" frameborder="0" allow="autoplay; encrypted-media" allowfullscreen></iframe>

---

## ⚙️ Component

You must have these installed first. to install the NONVIF project

### Software

- Ubuntu 20.04.5 LTS
  - Installation method ---> [Link](https://youtu.be/QKn5U2esuRk)
- Golang | version 1.19
  - Installation method ---> [Link](https://www.dropbox.com/scl/fi/9clhqzmcuk0gcjgby2h38/Golang.paper?dl=0&rlkey=pipvjgui4256g0wn4sp2aoklj)

### Hardware

- At least 2 IP cameras
  ##### or
- Use the "IP Webcam" application installed on your Android phone. | version 1.16.6.783
  - Installation method ---> [Link](https://play.google.com/store/apps/details?id=com.pas.webcam&hl=th)

### Dev Tool

- VScode or Visual Studio Code version | 1.76
  - Installation method ---> [Link](https://code.visualstudio.com/)

---

## 🔧 Installation

Install the `NONVIF Project` with the following command.

### Step : 1

Please clone this repository from Nuntakorn-Buu/NONVIF on Github.

```bash
  git clone https://github.com/Nuntakorn-Buu/NONVIF.git
```

### Step : 2

After completing the repository clone, use the following command to navigate into the Folder NONVIF

```bash
  cd NONVIF
```

### Step : 3

After that, use the command `< go get >` to install and update the necessary modules for this project.

```bash
  go get
```

---

## 💡 Running Tests

To run tests, Run the following command

```zsh
  go run .
```

If the `< go run . >` command is successful, you will see the following message below.

![go run .](https://github.com/Nuntakorn-Buu/NONVIF/blob/main/view/front-end/asset/images/go%20run.png?raw=true)

---

## 📑 Tech Stack

#### Backend

- HTTP Middleware ( [gorilla/mux](https://github.com/gorilla/mux) )
- Login Session ( [gorilla/sessions](https://github.com/gorilla/sessions) )
- Database Sqlite ( [go-sqlite3](https://github.com/mattn/go-sqlite3) & [msql](https://github.com/mateors/msql) )
- Package ONVIF ( [inspii/onvif](https://github.com/inspii/onvif) ) or ( [Nuntakorn-Buu/onvif-2](https://github.com/Nuntakorn-Buu/onvif-2) )

#### Frontend

- Server side templating ( [Go HTML Templates](https://pkg.go.dev/html/template) )
- Frontend `Pure ( HTML , CSS , JavaScript)`

#### Framework

- w3.css
  - Installation method ---> [Link](https://www.w3schools.com/w3css/)
- Font-Awesome
  - Home page ---> [Link](https://fontawesome.com/)
  - Installation method ---> [Link](https://cdnjs.com/libraries/font-awesome)

#### Modules

```go
require (
    github.com/gorilla/mux v1.8.0
	github.com/gorilla/sessions v1.2.1
	github.com/inspii/onvif v0.0.0-20220209003952-107e7e0b00d2
	github.com/mateors/msql v0.0.0-20230316154058-0966f727302f
	github.com/mattn/go-sqlite3 v1.14.17
)

// indirect
require (
    github.com/gorilla/securecookie v1.1.1
)
```

---

## 📧 Support & Feedback

If there is a problem and need help or If you have any feedback, you can contact us at

![Support email](https://img.shields.io/badge/Support-Email-green?style=plastic&logo=appveyor) - Send email to 62050553@go.buu.ac.th

![Support github](https://img.shields.io/badge/Support-Github-blue?style=plastic&logo=appveyor) - Create issues on https://github.com/Nuntakorn-Buu/NONVIF

---

## 👷🏻‍♂️ Used By

This project has been tested by :

- 62050553 Nuntakorn Sopap (Student), Faculty of Engineering, Burapha University, Department of Embedded Systems.

---

## 🙋🏻‍♂️ Author

- [@Nuntakorn [Non]](https://www.github.com/Nuntakorn-Buu)

---
