package main

import (
	"crypto/tls"
	"fmt"

	gomail "gopkg.in/mail.v2"
)

func main() {
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", "sender-email@gmail.com")

	// Set E-Mail receivers
	m.SetHeader("To", "remitent-email@gmail.com")

	// Set E-Mail subject
	m.SetHeader("Subject", "Gomail test subject")

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", "Hello there, this is an example. If you can read this, congratulations! You made it.")

	// Settings for SMTP server
	d := gomail.NewDialer("smtp.gmail.com", 587, "sender-email@gmail.com", "My Complicated Password")

	// This is only needed when SSL/TLS certificate is not valid on server.
	// For production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}

	return
}
