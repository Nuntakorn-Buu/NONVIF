package main

import (
	"fmt"
	"net/smtp"
)

var FromEmail = "your.email@gmail.com"                                       // แก้ไขเป็นที่อยู่อีเมล์ Gmail ของเรา
var EmailPassword = "รหัสที่ได้รับจากการทำการยืนยันขั้นสูง gmxyuycyfaaeywwe" // แก้ไขเป็นรหัสผ่านของบัญชี Gmail ของเรา
// หากเป็นบัญชีที่ใช้งาน"การยืนยันขั้นสูง" (2-Step Verification)
// ให้หารหัสจาก Link = https://nuntakorn-sp.notion.site/Problem-Send-Email-cb36c73588ca40de89d997e06581f017

func emailSend(email string) {
	from := FromEmail
	password := EmailPassword

	// Receiver email address.
	to := []string{email}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	subject := "Subject: GolanLoginSystem Account Recovery\n"

	mainMessage := fmt.Sprintf("<body>Welcome to master Academy, Your password  verification code is <h2 style=\"text-align:center;\"><span style=\"font-size:40px;border:2px solid black;padding:10px\">%v</span></h2> \n</body>", randomNUM)

	body := mainMessage
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	message := []byte(subject + mime + body)

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent : Successfully !!")
	fmt.Println("-------------------- | Please check your email.")

}
