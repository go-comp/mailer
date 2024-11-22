package mailer

type MailConfig struct {
	Transport   string
	Host        string
	Port        int
	Username    string
	Password    string
	Encryption  string
	Timeout     int // optional timeout in seconds
	LocalDomain string // for EHLO domain
	Region      string
}