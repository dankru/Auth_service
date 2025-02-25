package server

import (
	"fmt"
	"github.com/dankru/Auth_service/internal/service/auth"
	authpb "github.com/dankru/proto-definitions/pkg/auth"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	addr       string
	network    string
	grpcServer *grpc.Server
	authServer *auth.Server
}

func NewServer(network, addr string, authServer *auth.Server) *Server {
	grpcServer := grpc.NewServer()
	return &Server{
		addr:       addr,
		network:    network,
		grpcServer: grpcServer,
		authServer: authServer,
	}
}

func (s *Server) Run() error {
	lis, err := net.Listen(s.network, s.addr)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	authpb.RegisterTokenServiceServer(s.grpcServer, s.authServer)
	log.Printf("serving on %s", s.addr)

	if err := s.grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %w", err)
	}

	return nil
}
