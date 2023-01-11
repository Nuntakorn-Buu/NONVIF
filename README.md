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
- Dashboard page
- Setting page
- Add camera
- Remove camera
- Forget Password
---
# Stack
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
    github.com/mateors/msql v0.0.0-20211213034720-26a7ed4e79c4
    github.com/mattn/go-sqlite3 v1.14.11
)
```