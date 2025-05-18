package main

import (
	"fmt"
	"log"
	"net/smtp"
)

func main() {
	from := "habeebamoo08@gmail.com"
	password := "exphqvkpdzrbhrdp"

	to := []string{"alexjohnson99.uk@gmail.com"}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)

	msg := []byte("Subject: GoMail")

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, msg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Email sent successfully")
}