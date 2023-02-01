# NONVIF
* การประยุกต์ใช้โปรโตคอล ONVIF เพื่อรองรับการจัดการกล้องไอพีและการสตรีมวิดีโอผ่านเว็บไซต์ .
* Application of the ONVIF Protocol to support IP camera management and video streaming through websites .

# NONVIF System
## ONVIF Profile T 
- Streaming Video
- Metadata
---
## Features Projects
- Signup page
- Login page
- Logout page
- Home page
- Cameras page
- Dashboard page
- About page
- Portfolio page
- Add camera
- Remove camera
- Forget Password
---
# Explain Details
## Things to install before use
- Ubuntu 20.04.5 LTS
-   Installation method ---> [<Link>](https://youtu.be/QKn5U2esuRk)
- Golang version 1.19
-   Installation method ---> [<Link>](https://www.dropbox.com/scl/fi/9clhqzmcuk0gcjgby2h38/Golang.paper?dl=0&rlkey=pipvjgui4256g0wn4sp2aoklj)
## Backend
- HTTP Middleware [gorilla/mux](https://github.com/gorilla/mux)
- Login Session [gorilla/sessions](https://github.com/gorilla/sessions)
- Database Sqlite [go-sqlite3](https://github.com/mattn/go-sqlite3) & [msql](https://github.com/mateors/msql)
## Frontend
- Server side templating [Go HTML Templates](https://pkg.go.dev/html/template)
- Frontend ```Pure ( HTML , CSS )```
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