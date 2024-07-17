package main

import (
	"awesomeProject/Repo"
	"awesomeProject/Usecase"
	"awesomeProject/config"
	"awesomeProject/database"
	"awesomeProject/observability/logs"
	"awesomeProject/server"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

func main() {
	logs.InitLogger()
	defer func(LogFile *os.File) {
		err := LogFile.Close()
		if err != nil {
		}
	}(logs.LogFile)
	fields := logrus.Fields{"module": "main", "function": "main"}
	logs.LogInfo("Service started", fields)

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading config:", err)
	}

	db, err := database.InitDatabase(cfg)
	if err != nil {
		panic("failed to connect database")
	}
	Repository := Repo.NewGormRepo(db.SQL)
	UseCase := Usecase.NewUseCase(Repository)

	s := server.NewFiberServer()
	if err := s.Start(cfg.PORT, UseCase); err != nil {
		log.Fatal(err)
	}
}
