package main

import (
	"github.com/dankru/Auth_service/internal/service/auth"
	authpb "github.com/dankru/proto-definitions/pkg/auth"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %s", err.Error())
	}

	s := auth.NewAuthServer([]byte(os.Getenv("HMAC_SECRET")))
	grpcServer := grpc.NewServer()

	authpb.RegisterTokenServiceServer(grpcServer, s)
	log.Printf("serving")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err.Error())
	}
}
