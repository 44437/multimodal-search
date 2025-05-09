package main

import (
	"log"
	"multimodal-search/internal/config"
	"multimodal-search/internal/server"
	"multimodal-search/internal/multimodal"
	"os"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	zap.ReplaceGlobals(logger)

	appEnv := os.Getenv("APP_ENV")
	conf, err := config.New(".config", appEnv)
	if err != nil {
		log.Fatalln(err)
	}
	conf.Print()

	handler := multimodal.NewHandler(multimodal.NewService(multimodal.NewRepository(conf.Weaviate)))

	handlers := []server.Handler{handler}

	srv := server.NewServer(conf.Server.Port, handlers)

	err = srv.Start()
	if err != nil {
		log.Fatalln(err)
	}
}
