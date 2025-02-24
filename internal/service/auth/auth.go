package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/dankru/Auth_service/internal/domain"
	authpb "github.com/dankru/proto-definitions/pkg/auth"
	"github.com/golang-jwt/jwt"
	"math/rand"
	"strconv"
	"time"
)

type Server struct {
	authpb.UnimplementedTokenServiceServer
	hmacSecret         []byte
	sessionsRepository SessionsRepository
}

type SessionsRepository interface {
	Create(token domain.RefreshSession) error
	Get(token string) (domain.RefreshSession, error)
}

func NewAuthServer(hmac []byte) *Server {
	return &Server{
		UnimplementedTokenServiceServer: authpb.UnimplementedTokenServiceServer{},
		hmacSecret:                      hmac,
		sessionsRepository:              nil,
	}
}

func (s *Server) GenerateToken(ctx context.Context, userData *authpb.UserData) (*authpb.JWT, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Subject:   userData.Id,
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Hour * 15).Unix(),
	})

	accessToken, err := t.SignedString(s.hmacSecret)
	if err != nil {
		return &authpb.JWT{
			AccessToken:  "",
			RefreshToken: "",
		}, err
	}

	refreshToken, err := newRefreshToken()
	if err != nil {
		return &authpb.JWT{
			AccessToken:  "",
			RefreshToken: "",
		}, err
	}

	userId, err := strconv.ParseInt(userData.Id, 10, 64)
	if err != nil {
		return &authpb.JWT{
			AccessToken:  "",
			RefreshToken: "",
		}, err
	}
	if err := s.sessionsRepository.Create(domain.RefreshSession{
		UserID:    userId,
		Token:     refreshToken,
		ExpiresAt: time.Now().Add(time.Hour * 24 * 30),
	}); err != nil {
		return &authpb.JWT{
			AccessToken:  "",
			RefreshToken: "",
		}, err
	}

	return &authpb.JWT{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, err
}

func (s *Server) ParseToken(ctx context.Context, token *authpb.TokenRequest) (*authpb.UserData, error) {
	t, err := jwt.Parse(token.Token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return s.hmacSecret, nil
	})

	if err != nil {
		return &authpb.UserData{Id: "0"}, err
	}

	if !t.Valid {
		return &authpb.UserData{Id: "0"}, errors.New("invalid token")
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return &authpb.UserData{Id: "0"}, errors.New("invalid claims")
	}

	subject, ok := claims["sub"].(string)
	if !ok {
		return &authpb.UserData{Id: "0"}, errors.New("Invalid subject")
	}

	id, err := strconv.Atoi(subject)
	if err != nil {
		return &authpb.UserData{Id: "0"}, errors.New("invalid subject")
	}

	return &authpb.UserData{Id: strconv.Itoa(id)}, nil
}

func newRefreshToken() (string, error) {
	b := make([]byte, 32)

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	if _, err := r.Read(b); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", b), nil
}
