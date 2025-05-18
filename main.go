package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/smtp"
	"os"
)

type Config struct {
	Sender string `json:"sender"`
	Password string `json:"password"`
}

func readMessage() ([]byte, error) {
	msg, err := os.ReadFile("res/message.txt")
	if err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir("res", 0755)
			check(err)

			_, err = os.Create("res/message.txt")
			check(err)
		} else {
			return nil, errors.New("can't read res/message.txt file")
		}
	}

	return msg, nil
}

func saveCred(sender, password string) error {
	config := Config{sender, password}
	configJson, err := json.MarshalIndent(config, "", " ")
	check(err)

	err = os.WriteFile("config.json", configJson, 0644)
	if err != nil {
		if os.IsNotExist(err) {
			_, err = os.Create("config.json")
			check(err)
		} else {
			return errors.New("can't read config.json file")
		}
	}

	fmt.Println("Your credentials have been saved")
	return nil
}

func getCred() (Config, error) {
	configJson, err := os.ReadFile("config.json")
	if err != nil {
		return Config{}, errors.New("can't read config file")
	}

	var data Config
	err = json.Unmarshal(configJson, &data)
	if err != nil {
		return Config{}, errors.New("can't format config file")
	}

	return data, nil
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	var action string
	var sender string
	var pass string
	var receiver string

	flag.StringVar(&action, "a", "", "save or send")
	flag.StringVar(&sender, "from", "", "Senders Gmail's Address")
	flag.StringVar(&pass, "password", "", "Senders Gmail's App Password")
	flag.StringVar(&receiver, "to", "", "Receiver's Email Address")

	flag.Parse()

	switch action {
	case "save":
		if action == "" || sender == "" || pass == "" {
			fmt.Println("You must provide action, sender and password")
			return
		}

		err := saveCred(sender, pass)
		check(err)
	case "send":
		if action == "" || receiver == "" {
			fmt.Println("Define you action -a send")
			return
		}

		data, err := getCred()
		check(err)
		
		from := data.Sender
		password := data.Password

		to := []string{receiver}

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
}