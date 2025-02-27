package main

import (
	"github.com/dankru/Auth_service/internal/server"
	"github.com/dankru/Auth_service/internal/service/auth"
	"log"
	"os"
)

func main() {
	authServer := auth.NewAuthServer([]byte(os.Getenv("HMAC_SECRET")))

	srv := server.NewServer("tcp", ":8081", authServer)
	if err := srv.Run(); err != nil {
		log.Fatal("failed to run: %s", err.Error())
	}
}
