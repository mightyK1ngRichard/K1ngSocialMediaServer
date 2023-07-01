package app

import (
	"K1ngSochialMediaServer/internal/app/config"
	"K1ngSochialMediaServer/internal/app/repository"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Application struct {
	Config     *config.Config
	Logger     *logrus.Logger
	Router     *gin.Engine
	Repository *repository.Repository
}

func New(c *config.Config, l *logrus.Logger, r *gin.Engine, rep *repository.Repository) *Application {
	return &Application{
		Config:     c,
		Router:     r,
		Logger:     l,
		Repository: rep,
	}
}

func (a *Application) Run() {
	defer func() {
		if err := a.Repository.TurnOffDataBase(); err != nil {
			a.Logger.Fatalln(err)
		}
		a.Logger.Info("Data base was turned off")
	}()

	a.Logger.Info("Server start up")
	a.Router.Static("/static/img", "./static/img")

	addr := fmt.Sprintf(":%d", a.Config.ServicePort)
	if err := a.Router.Run(addr); err != nil {
		a.Logger.Fatalln(err)
	}
	a.Logger.Info("Server down")
}
