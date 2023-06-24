package main

import (
	"K1ngSochialMediaServer/internal/app/config"
	"K1ngSochialMediaServer/internal/app/dsn"
	"K1ngSochialMediaServer/internal/app/handler"
	app "K1ngSochialMediaServer/internal/app/pkg"
	"K1ngSochialMediaServer/internal/app/repository"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"log"
)

func main() {
	log.Println("Application start!")
	router := gin.Default()
	logger := logrus.New()
	postgresqlLine := dsn.FromEnv(logger)
	if postgresqlLine == "" {
		logger.Error("Postgres line is empty")
	}

	rep, err := repository.NewRepository(postgresqlLine, logger)
	if err != nil {
		log.Fatalln(err)
	}

	conf, err := config.NewConfig(logger)
	if err != nil {
		log.Fatalln(err)
	}

	a := app.New(conf, logger, router, rep)
	hand := handler.NewHandler(a)
	hand.Register(a.Router)

	a.Run()
	log.Println("Application terminated!")
}
