// Package smtp defines email adaptor
package smtp

import (
	"context"
	"fmt"
	"text/template"

	"github.com/ezex-io/ezex-notification/internal/ports"
	"gopkg.in/gomail.v2"
)

type SMTPAdapter struct {
	config        *Config
	templateCache map[string]*template.Template
	dialer        *gomail.Dialer
}

func NewSMTPAdapter(confs *Config) ports.EmailPort {
	dialer := gomail.NewDialer(
		confs.Host,
		confs.Port,
		confs.User,
		confs.Pass,
	)

	return &SMTPAdapter{
		dialer:        dialer,
		config:        confs,
		templateCache: make(map[string]*template.Template),
	}
}

func (a *SMTPAdapter) SendEmail(_ context.Context, recipient, subject, body string) error {
	message := gomail.NewMessage()
	message.SetHeader("From", a.config.FromEmail)
	message.SetHeader("To", recipient)
	message.SetHeader("Subject", subject)

	// Add HTML alternative
	message.AddAlternative("text/html", body)

	if err := a.dialer.DialAndSend(message); err != nil {
		return fmt.Errorf("failed to send email to %s: %w", recipient, err)
	}

	return nil
}
