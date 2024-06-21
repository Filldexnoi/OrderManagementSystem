package main

import (
	"awesomeProject/Repo"
	"awesomeProject/Usecase"
	"awesomeProject/config"
	"awesomeProject/database"
	"awesomeProject/server"
	"log"
)

func main() {
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
	s.SetupFiberRoute(UseCase)
	if err := s.Start(cfg.PORT); err != nil {
		log.Fatal(err)
	}
}
