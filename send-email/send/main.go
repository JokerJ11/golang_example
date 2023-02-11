package main

import (
	"log"

	"net/smtp"
)

func main() {

	// Mailtrap account config

	// username := "john"
	from := "kyawyeaungcu48@gmail.com"

	password := "nxhrordwjsqduckh"

	smtpHost := "smtp.gmail.com"

	// Choose auth method and set it up

	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Message data

	to := []string{
		"kyawyeaung.cu@gmail.com",
		"skkosmith@gmail.com",
	}

	message := []byte(`To: kyawyeaung.cu@gmail.com

	From: kyawyeaungcu48@gmail.com

	Subject: Why aren’t you using Mailtrap yet?

	Here’s the space for your great sales pitch`)

	// Connect to the server and send message

	smtpUrl := smtpHost + ":25"

	err := smtp.SendMail(smtpUrl, auth, from, to, message)

	if err != nil {

		log.Fatal(err)

	}
}
