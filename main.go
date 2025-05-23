package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"gopkg.in/gomail.v2"
)

type Config struct {
	Sender  string  `json:"sender"`
	Password  string  `json:"password"`
}

var (
	ErrReadingMsgFile = errors.New("can't read res/message.txt file")
	ErrReadingConfigFile = errors.New("can't read config.json file")
	ErrFormatingConfigFile = errors.New("can't format config file")
)

func readMessage() ([]byte, error) {
	msg, err := os.ReadFile("res/message.txt")
	if err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir("res", 0755)
			check(err)

			_, err = os.Create("res/message.txt")
			check(err)
			
			return []byte("Hello"), nil
		} else {
			return nil, ErrReadingMsgFile
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
			return ErrReadingConfigFile
		}
	}

	fmt.Println("Your credentials have been saved")
	return nil
}

func getCred() (Config, error) {
	configJson, err := os.ReadFile("config.json")
	if err != nil {
		return Config{}, ErrReadingConfigFile
	}

	var user Config
	err = json.Unmarshal(configJson, &user)
	if err != nil {
		return Config{}, ErrFormatingConfigFile
	}

	return user, nil
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

		user, err := getCred()
		check(err)

		to := strings.Split(receiver, ",")
		 for i := range to {
			to[i] = strings.TrimSpace(to[i])
		}

		body, err := readMessage()
		check(err)

		m := gomail.NewMessage()
		m.SetHeader("From", user.Sender)
		m.SetHeader("To", to...)
		m.SetHeader("Subject", "GoMail")
		m.SetBody("text/plain", string(body))
		
		d := gomail.NewDialer("smtp.gmail.com", 465, user.Sender, user.Password)
		d.SSL = true

		if err := d.DialAndSend(m); err != nil {
			fmt.Println("Couldn't send Email")
			log.Fatal(err)
		}

		fmt.Println("Email sent successfully")
	default:
		fmt.Println("Command not supported")
	}
}