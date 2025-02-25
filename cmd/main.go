package main

import (
	"github.com/dankru/Auth_service/internal/server"
	"github.com/dankru/Auth_service/internal/service/auth"
	"log"
)

func main() {
	authServer := auth.NewAuthServer([]byte("eqweqwenqwr"))

	srv := server.NewServer("tcp", ":9000", authServer)
	if err := srv.Run(); err != nil {
		log.Fatal("failed to run: %s", err.Error())
	}
}
