package grpc

import (
	"context"
	"fmt"

	"github.com/ezex-io/ezex-notification/api/grpc/proto"
	"github.com/ezex-io/ezex-notification/internal/interactors"
)

type NotificationService struct {
	proto.UnimplementedNotificationServiceServer
	emailWorker *interactors.EmailWorker
}

func NewNotificationService(emailWorker *interactors.EmailWorker) *NotificationService {
	n := &NotificationService{
		emailWorker: emailWorker,
	}

	return n
}

func (ns *NotificationService) SendEmail(
	ctx context.Context,
	req *proto.SendEmailRequest,
) (*proto.SendEmailResponse, error) {
	_, err := ns.emailWorker.SendEmail(ctx, req.Recipient, req.Subject, req.TemplateName, req.TemplateFields)
	if err != nil {
		return &proto.SendEmailResponse{}, fmt.Errorf("failed to send email: %w", err)
	}

	return &proto.SendEmailResponse{}, nil
}
