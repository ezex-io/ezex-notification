// Package grpc defines it protocol server
package grpc

import (
	"fmt"
	"net"

	"github.com/ezex-io/ezex-notification/api/grpc/proto"
	"google.golang.org/grpc"
)

type Server struct {
	sv      *grpc.Server
	ln      net.Listener
	errCh   chan error
	configs Config
}

func NewServer(notificationService *NotificationService, conf Config) (*Server, error) {
	listener, err := net.Listen("tcp", ":"+conf.Port)
	if err != nil {
		return nil, fmt.Errorf("failed to listen on port: %w", err)
	}

	srv := grpc.NewServer()
	proto.RegisterNotificationServiceServer(srv, notificationService)

	return &Server{
		sv:      srv,
		ln:      listener,
		errCh:   make(chan error),
		configs: conf,
	}, nil
}

func (s *Server) Start() {
	s.errCh <- s.sv.Serve(s.ln)
}

func (s *Server) Notify() <-chan error {
	return s.errCh
}
