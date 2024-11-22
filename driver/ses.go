package driver

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/smithy-go/middleware"
)

type SESMailer struct {
	client    *ses.Client
	sender    string
	region    string
}

func NewSESMailer(sender string, region string) (*SESMailer, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		return nil, fmt.Errorf("unable to load SDK config, %v", err)
	}

	client := ses.NewFromConfig(cfg)
	return &SESMailer{
		client: client,
		sender: sender,
		region: region,
	}, nil
}

func (m *SESMailer) Send(to string, subject string, body string) error {
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []string{to},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Text: &ses.Content{
					Data:    aws.String(body),
					Charset: aws.String("UTF-8"),
				},
			},
			Subject: &ses.Content{
				Data:    aws.String(subject),
				Charset: aws.String("UTF-8"),
			},
		},
		Source: aws.String(m.sender),
	}

	result, err := m.client.SendEmail(context.TODO(), input)
	if err != nil {
		return fmt.Errorf("failed to send email, %v", err)
	}

	fmt.Printf("Email sent successfully to %s with Message ID: %s\n", to, *result.MessageId)
	return nil
}
