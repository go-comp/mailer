package mailer

type PostmarkMailer struct {
	config *MailConfig
}

func NewPostmarkMailer(config *MailConfig) *PostmarkMailer {
	return &PostmarkMailer{config: config}
}

func (m *PostmarkMailer) Send(to string, subject string, body string) error {
	// Implement Postmark sending logic here.
	return nil // Placeholder return.
}
