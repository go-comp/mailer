package mailer

import (
	"github.com/stretchr/testify/mock"
)

// MockMailer is a mock implementation of the Mailer interface
type MockMailer struct {
	mock.Mock
}

func (m *MockMailer) Send(to string, subject string, body string) error {
	args := m.Called(to, subject, body)
	return args.Error(0)
}
