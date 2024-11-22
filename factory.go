package mailer

import (
	"fmt"
	"github.com/go-comp/mailer/driver"
)

type MailFactory struct{}

func (f *MailFactory) CreateMailer(config *MailConfig) (Mailer, error) {
	switch config.Transport {
	case "smtp":
		return NewSMTPMailer(config), nil
	case "ses":
		return NewSESMailer(config.Username, config.Region) // Assuming Username is used as sender email
	case "postmark":
		return NewPostmarkMailer(config), nil
	case "resend":
		return NewResendMailer(config), nil
	default:
		return nil, fmt.Errorf("unsupported transport: %s", config.Transport)
	}
}