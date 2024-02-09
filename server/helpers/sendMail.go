package helpers

import (
	"fmt"
	"net/smtp"

	"github.com/kawojue/go-initenv"
)

func SendMail(to []string, message string) {
	email := initenv.GetEnv("EMAIL", "")
	email_pswd := initenv.GetEnv("EMAIL_PSWD", "")

	auth := smtp.PlainAuth(
		"",
		email,
		email_pswd,
		"smtp.gmail.com",
	)

	if err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		email,
		to,
		[]byte(message),
	); err != nil {
		panic(err)
	}

	fmt.Println("Email sent successfully.")
}
