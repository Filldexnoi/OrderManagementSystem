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
	UseCase := new(Usecase.UseCase)
	productRepo := Repo.NewProductRepo(db.SQL)
	UseCase.Product = Usecase.NewProductUseCase(productRepo)
	stockRepo := Repo.NewStock(db.SQL)
	UseCase.Stock = Usecase.NewStockUseCase(stockRepo)
	transactionRepo := Repo.NewTransactionRepo(db.SQL)
	UseCase.Transaction = Usecase.NewTransactionUseCase(transactionRepo, productRepo)
	orderRepo := Repo.NewOrderRepo(db.SQL)
	UseCase.Order = Usecase.NewOrderUseCase(orderRepo)
	s := server.NewFiberServer()
	s.SetupFiberRoute(UseCase)
	if err := s.Start(cfg.PORT); err != nil {
		log.Fatal(err)
	}
}
