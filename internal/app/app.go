package app

import (
	"log"
	"net/http"
	"waifunet/internal/config"
	database "waifunet/internal/database/postgresql"
	"waifunet/internal/transport/rest"

	"github.com/sirupsen/logrus"
)

type Application struct {
	config *config.Config
	logger *logrus.Logger
	router *http.ServeMux
}

func NewApp(cfg *config.Config, logger *logrus.Logger) Application {
	return Application{
		config: config.GetConfig(),
		logger: logrus.New(),
		router: http.NewServeMux(),
	}
}

func (a *Application) Start() error {
	return a.startHttp()
}

func (a *Application) startHttp() error {
	// инициализация роутера
	rest.CreateRoutes(a.router)

	// инициализация дб - создание соединения с бд
	err := database.PostgreRoutine(a.config)
	if err != nil {
		log.Println("err creating conn pool")
	}

	// запуск http-сервера
	log.Println("server starting at localhost:5445")
	return http.ListenAndServe(":5445", a.router)
}
