package main

import (
	"log"
	"strings"

	"github.com/emersion/go-sasl"
	"github.com/emersion/go-smtp"
)

func SendMail() {
	// Set up authentication information.
	auth := sasl.NewPlainClient("", "user@example.com", "password")

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{"recipient@example.net"}
	msg := strings.NewReader("To: recipient@example.net\r\n" +
		"Subject: discount Gophers!\r\n" +
		"\r\n" +
		"This is the email body.\r\n")
	err := smtp.SendMail("169.24.2.82:1025", auth, "sender@example.org", to, msg)
	if err != nil {
		log.Fatal(err)
	}
	msg2 := strings.NewReader("To: recipient@example.net\r\n" +
		"Subject: discount Gophers!\r\n" +
		"\r\n" +
		"This is the email body.\r\n")
	log.Println(msg)
	err = smtp.SendMail("169.24.2.82:1025", auth, "sender2@example.org", to, msg2)

	if err != nil {
		log.Fatal(err)
	}
}
