package service

import (
	"fmt"
	"net/smtp"
	"strings"
	"sync"
)

type SmtpServiceInterface interface{}

type smtpService struct {
}

func SmtpService() *smtpService {
	return &smtpService{}
}

const CONFIG_SMTP_HOST = "10.54.59.13"
const CONFIG_SMTP_PORT = 25004
const CONFIG_SENDER_NAME = "PT. Makmur Subur Jaya <emailanda@gmail.com>"
const CONFIG_AUTH_EMAIL = "mygrapari@telkomsel.com"
const CONFIG_AUTH_PASSWORD = ""

func (smtpService *smtpService) SendEmail(wg *sync.WaitGroup) error {
	defer wg.Done()
	to := []string{"uci3026@gmail.com"}
	cc := []string{"adiamahardika.work@gmail.com"}
	subject := "Test mail"
	message := "Hello"
	body := "From: " + CONFIG_SENDER_NAME + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Cc: " + strings.Join(cc, ",") + "\n" +
		"Subject: " + subject + "\n\n" +
		message

	auth := smtp.PlainAuth("", CONFIG_AUTH_EMAIL, CONFIG_AUTH_PASSWORD, CONFIG_SMTP_HOST)
	smtpAddr := fmt.Sprintf("%s:%d", CONFIG_SMTP_HOST, CONFIG_SMTP_PORT)

	err := smtp.SendMail(smtpAddr, auth, CONFIG_AUTH_EMAIL, append(to, cc...), []byte(body))
	if err != nil {
		return err
	}

	return nil
}
