package main

import (
	"log"
	"waifunet/internal/app"
	"waifunet/internal/config"

	"github.com/sirupsen/logrus"
)

func main() { // (test)
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	log.Print("config initializing")
	cfg := config.GetConfig()

	log.Print("logger initializing")
	logger := logrus.New()

	app := app.NewApp(cfg, logger)
	return app.Start()
}
