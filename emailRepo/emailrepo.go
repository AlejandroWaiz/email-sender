package emailrepo

import (
	"fmt"
	"net/smtp"
	"os"
)

var auth smtp.Auth
var message []byte
var to []string
var smtpPort string
var from string

//CreateMail Create an email
func CreateMail() {

	// Sender data.
	from = "alejandronicolas.waiz@gmail.com"
	password := "olakease123"

	// Receiver email address.
	to = []string{
		"alejandronicolas.waiz@gmail.com",
	}

	// smtp server configuration.
	smtpHost := os.Getenv("smtpHost")
	smtpPort = os.Getenv("smtpPort")

	// Message.
	message = []byte("This is a test email message.")

	// Authentication.
	auth = smtp.PlainAuth("", from, password, smtpHost)

	fmt.Println(smtpHost, smtpPort)

}

//GetAuth returns auth
func GetAuth() smtp.Auth {

	return auth

}

//GetMessage returns message
func GetMessage() []byte {

	return message

}

//GetTo returns the target of the email
func GetTo() []string {

	return to

}

//GetFrom returns the email of the sender
func GetFrom() string {

	return from

}
