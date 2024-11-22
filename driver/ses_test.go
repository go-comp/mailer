package mailer

import (
    "context"
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
)

// MockSESClient is a mock implementation of the SES client.
type MockSESClient struct {
    mock.Mock
}

func (m *MockSESClient) SendEmail(ctx context.Context, params *ses.SendEmailInput, optFns ...func(*ses.Options)) (*ses.SendEmailOutput, error) {
    args := m.Called(ctx, params)
    return args.Get(0).(*ses.SendEmailOutput), args.Error(1)
}

func TestSESMailerSend(t *testing.T) {
    mockClient := new(MockSESClient)
    mockMailer := &SESMailer{client: mockClient, sender: "sender@example.com"}

    mockClient.On("SendEmail", mock.Anything, mock.Anything).Return(&ses.SendEmailOutput{MessageId: aws.String("1234")}, nil)

    err := mockMailer.Send("recipient@example.com", "Test Subject", "This is a test email.")

    assert.NoError(t, err)
    mockClient.AssertExpectations(t)
}
