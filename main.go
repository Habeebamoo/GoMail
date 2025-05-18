package main

import (
	"errors"
	"fmt"
	"log"
	"net/smtp"
	"os"
)

func readMessage() ([]byte, error) {
	msg, err := os.ReadFile("res/message.txt")
	if err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir("res", 0755)
			check(err)

			_, err = os.Create("res/message.txt")
			check(err)
		}
		return nil, errors.New("can't read res/message.txt file")
	}

	return msg, nil
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	from := "habeebamoo08@gmail.com"
	password := "exphqvkpdzrbhrdp"

	to := []string{"alexjohnson99.uk@gmail.com"}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)

	body, err := readMessage()
	check(err)
	msg := []byte("Subject: GoMail\r\n\r\n" + string(body))

	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, msg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Email sent successfully")
}