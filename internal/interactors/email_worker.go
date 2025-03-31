// Package interactors defines business logic
package interactors

import (
	"context"
	"fmt"

	"github.com/ezex-io/ezex-notification/internal/ports"
	"github.com/ezex-io/ezex-notification/templates"
)

type EmailWorker struct {
	emailPort       ports.EmailPort
	templateManager templates.TemplateManager
}

func NewEmailWorker(emailPort ports.EmailPort) *EmailWorker {
	tmpManager := templates.New()

	return &EmailWorker{
		emailPort:       emailPort,
		templateManager: *tmpManager,
	}
}

func (s *EmailWorker) SendEmail(
	ctx context.Context,
	recipient string,
	subject string,
	templateName string,
	templateFields map[string]string,
) (string, error) {
	body, err := s.templateManager.Render(templateName, templateFields)
	if err != nil {
		return "", fmt.Errorf("failed to render: %w", err)
	}

	if err := s.emailPort.SendEmail(ctx, recipient, subject, body); err != nil {
		return "", err
	}

	return body, nil
}
