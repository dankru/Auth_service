package main

import (
	"github.com/dankru/Auth_service/internal/repository/pg_repo"
	"github.com/dankru/Auth_service/internal/server"
	"github.com/dankru/Auth_service/internal/service/auth"
	"github.com/dankru/Commissions_simple/pkg/database/pg_db"

	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	conn := pg_db.Connection{
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_USER:     os.Getenv("DB_USER"),
		DB_NAME:     os.Getenv("DB_NAME"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
	}

	postgres := pg_db.NewPostgreSQLDB(conn)
	defer postgres.Close()

	repository := pg_repo.NewTokensRepository(postgres.DB)
	authServer := auth.NewAuthServer([]byte(os.Getenv("HMAC_SECRET")), repository)
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
