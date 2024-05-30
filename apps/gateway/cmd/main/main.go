package main

import (
	"log"
	"salesbot-kommo/apps/gateway/internal/queue"
	"salesbot-kommo/apps/gateway/internal/router"
	"salesbot-kommo/apps/gateway/internal/service"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Service service.Config
	Router  router.Config
	Queue   queue.Config
}

func main() {
	// Конфигурация приложения
	config := Config{}

	// Загрузка переменных окружения
	if err := cleanenv.ReadEnv(&config); err != nil {
		log.Fatalln(err)
	}

	Queue, err := queue.New(&config.Queue)
	if err != nil {
		log.Fatalln(err)
	}

	// Инициализация сервиса
	// код сервиса /internal/service/service.go
	Service, err := service.New(&config.Service, Queue)
	if err != nil {
		log.Fatalln(err)
	}

	// Инициализация HTTP-роутера
	// код роутера /internal/router/router.go
	Router, err := router.New(&config.Router, Service)
	if err != nil {
		log.Fatalln(err)
	}

	// Запуск роутера
	if err := Router.Listen(); err != nil {
		log.Fatalln(err)
	}
}
