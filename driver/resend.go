package driver

type ResendMailer struct {
	config *MailConfig
}

func NewResendMailer(config *MailConfig) *ResendMailer {
	return &ResendMailer{config: config}
}

func (m *ResendMailer) Send(to string, subject string, body string) error {
	// Implement Resend sending logic here.
	return nil // Placeholder return.
}
