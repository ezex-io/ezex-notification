// Package ports defines interfaces (ports)
package ports

import (
	"context"
)

type EmailPort interface {
	SendEmail(ctx context.Context, recipient string, subject string, body string) error
}
