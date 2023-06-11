// `package main` is the declaration of the package name.
// In Go, the `main` package is a special package that is used to build executable programs.
// It must contain a `main` function as the entry point of the program.
package main

// The `import` statement is used to import packages that are necessary for the program to run.
// In this case, the program is importing the following packages:
import (
	"fmt"      // Tools for manipulating the screen display. and reading data from the user You can use the Printf() function.
	"net/smtp" // It is used to send emails via the SMTP protocol (Simple Mail Transfer Protocol), which is the protocol used to send emails across networks.
)

var FromEmail = "your_email@gmail.com"         // Fixed to our Gmail email address.
var EmailPassword = "such as gmxyuycyfaaeywwe" // Codes obtained by performing advanced verification,
// If it is an account that uses "Advanced Verification" (2-Step Verification)
// How to create a code Link = https://nuntakorn-sp.notion.site/Problem-Send-Email-cb36c73588ca40de89d997e06581f017

// The function sends an email with a verification code to a specified email address using SMTP authentication.
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
