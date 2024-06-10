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

	productRepo := Repo.NewProductRepo(db.SQL)
	productUseCase := Usecase.NewProductUseCase(productRepo)

	stockRepo := Repo.NewStock(db.SQL)
	stockUseCase := Usecase.NewStockUseCase(stockRepo)
	s := server.NewFiberServer()
	s.SetupFiberRoute(productUseCase, stockUseCase)
	if err := s.Start(cfg.PORT); err != nil {
		log.Fatal(err)
	}
}
