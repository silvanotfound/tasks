package main

import (
	"silvanotfound/tasks/config"
	"silvanotfound/tasks/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.NewLoger("main")
	err := config.Init()
	if err != nil {
		logger.Errorf("Erro ao inicializaro o bando de dados", err.Error())
		return
	}
	router.Initialize()
}
