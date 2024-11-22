package driver

import (
	"fmt"
	"net/smtp"
)

type SMTPMailer struct {
	config *MailConfig
}

func NewSMTPMailer(config *MailConfig) *SMTPMailer {
	return &SMTPMailer{config: config}
}

func (m *SMTPMailer) Send(to string, subject string, body string) error {
	auth := smtp.PlainAuth("", m.config.Username, m.config.Password, m.config.Host)
	msg := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n")

	err := smtp.SendMail(fmt.Sprintf("%s:%d", m.config.Host, m.config.Port), auth, m.config.Username, []string{to}, msg)
	if err != nil {
		return err
	}
	return nil
}
