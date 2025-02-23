package auth

import (
	"context"
	authpb "github.com/dankru/proto-definitions/pkg/auth"
)

type Server struct {
	authpb.UnimplementedTokenServiceServer
}

func (s *Server) GenerateToken(context.Context, *authpb.UserData) (*authpb.JWT, error) {
	var jwt authpb.JWT
	return &jwt, nil
}

func (s *Server) ParseToken(context.Context, *authpb.TokenRequest) (*authpb.UserData, error) {
	var userData authpb.UserData
	userData.Id = "User id here, message successfully received!"
	return &userData, nil
}
