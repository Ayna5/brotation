package grpc

import (
	"context"
	"fmt"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/Ayna5/bannersRotation/internal/app"
	banners_rotation_pb "github.com/Ayna5/bannersRotation/pkg/banners-rotation"
)

type Server struct {
	banners_rotation_pb.UnimplementedBannersRotationServer
	lis    net.Listener
	l      *logrus.Logger
	server *grpc.Server
	api    *app.App
}

func NewServer(l *logrus.Logger, api *app.App, address string) (*Server, error) {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return nil, fmt.Errorf("start listen error: %w", err)
	}

	serverPayloadLoggingDecider := func(ctx context.Context, fullMethodName string, servingObject interface{}) bool { return true }
	server := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_logrus.UnaryServerInterceptor(&logrus.Entry{Logger: l}),
				grpc_logrus.PayloadUnaryServerInterceptor(&logrus.Entry{Logger: l}, serverPayloadLoggingDecider)),
		))

	srv := &Server{
		lis:    lis,
		l:      l,
		server: server,
		api:    api,
	}
	banners_rotation_pb.RegisterBannersRotationServer(server, srv)

	return srv, nil
}

func (s *Server) Start() error {
	if err := s.server.Serve(s.lis); err != nil {
		return fmt.Errorf("start server error: %w", err)
	}
	return nil
}

func (s *Server) Stop() error {
	if s.server == nil {
		return errors.New("grpc server is nil")
	}

	s.server.GracefulStop()
	return nil
}
