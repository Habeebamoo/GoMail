# GoMail
GoMail is a free and convienent CLI tool for sending emails

## Why
GoMail saves you the time of sending emails repetitively with your desktop app, instead you can compose an email to any amount of recipient in seconds

# Proceedure
1. Enable 2FA authentication on your Google account
2. Generate an App Password from your account
3. Use the application easily

# Example
### Go to [myaccount.google.com](https://myaccount.google.com) and Enable 2FA in your Google account
![screenshot](assets/first.png)

### Go to [account.google.com](https://myaccount.google.com/apppasswords), create an app and copy the password
![screenshot](assets/second.png)

## Setup
For Linux/MacOS
```bash
go build -o main
```

Run with 
```bash
./main
```

For Windows (It has been done for you already, No Need)
```bash
go build -o main.exe
```

Run with
```bash
main.exe
```
OR...
```bash
.\main.exe
```

# How to use
Firstly save your email address and the password gotten from Google with the command below, it will be added to the config.json file where the program will utilize it

command
```bash
.\main.exe -a save -from "youremail" -password "your-generated-app-password"
```

Secondly, you should go into the *message.txt* file where you compose your actual message. 

*directory: res > message.txt*

Finally send your composed message to your receiver's email address easily with the command below

```bash
.\main.exe -a send -to "receipent's email"
```

To send to multiple receipients, seperate them by a comma ","

```bash
.\main.exe -a send -to "first@gmail.com, second@gmail.com, third@gmail.com"
```

#### NB: you should compose your message in the ***message.txt*** in the ***res*** folder

