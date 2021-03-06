package emailserv

import (
	"fmt"
	"log"
	"net/smtp"

	"github.com/AlejandroWaiz/email-sender/emailrepo"
)

//SendMail sends the email
func SendMail() {

	emailrepo.CreateMail()
	auth := emailrepo.GetAuth()
	log.Println(auth)

	// Sending email.
	err := smtp.SendMail("smtp.gmail.com"+":587", auth,
		emailrepo.GetFrom(), emailrepo.GetTo(), emailrepo.GetMessage())

	if err != nil {
		fmt.Println("[SendMail]", err)
		return
	}
	fmt.Println("Email Sent Successfully!")

}
