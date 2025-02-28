package main

import (
	"github.com/dankru/Auth_service/internal/server"
	"github.com/dankru/Auth_service/internal/service/auth"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	authServer := auth.NewAuthServer([]byte(os.Getenv("HMAC_SECRET")))
	srv := server.NewServer("tcp", viper.GetString("server.port"), authServer)
	if err := srv.Run(); err != nil {
		log.Fatal("failed to run: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("../configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
