package mailer

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateSMTPMailer(t *testing.T) {
	config := &MailConfig{
		Transport: "smtp",
		Host:      "smtp.mailtrap.io",
		Port:      2525,
		Username:  "your_username",
		Password:  "your_password",
	}

	factory := &MailFactory{}
	mailer, err := factory.CreateMailer(config)

	assert.NoError(t, err)
	assert.IsType(t, &SMTPMailer{}, mailer)
}

func TestCreateUnsupportedMailer(t *testing.T) {
	config := &MailConfig{
		Transport: "unsupported",
	}

	factory := &MailFactory{}
	mailer, err := factory.CreateMailer(config)

	assert.Error(t, err)
	assert.Nil(t, mailer)
}
