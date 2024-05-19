package email

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendEmails(emails []string, rate float64) error {
	auth := authAccountInSmtp()
	msg := fmt.Sprintf("Subject: Зміни курсу BTC до UAH\nТепер біток коштує  %.1f грн", rate)

	err := smtp.SendMail(
		os.Getenv("SMTP_ADDR"),
		auth,
		os.Getenv("SMTP_NAME"),
		emails,
		[]byte(msg),
	)

	return err
}

func authAccountInSmtp() smtp.Auth {
	return smtp.PlainAuth(
		"",
		os.Getenv("SMTP_NAME"),
		os.Getenv("SMTP_PASS"),
		os.Getenv("SMTP_HOST"),
	)
}
