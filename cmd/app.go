package main

import (
	"context"
	"log"
	"net/http"

	"github.com/v1adhope/web-service-employees/internal/config"
	v1 "github.com/v1adhope/web-service-employees/internal/controller/http/v1"
	"github.com/v1adhope/web-service-employees/internal/usecase"
	"github.com/v1adhope/web-service-employees/internal/usecase/repository"
	"github.com/v1adhope/web-service-employees/pkg/mongodb"
)

const (
	_dbName  = "employeeStorage"
	_colName = "employee"
)

func main() {
	cfg := config.GetConfig()

	mongoClient, err := mongodb.New(context.Background(), cfg.ConStr)
	if err != nil {
		log.Fatal(err)
	}
	defer mongoClient.Close(context.Background())

	repo := repository.New(mongoClient.GetCollecion(_dbName, _colName))

	usecase := usecase.New(repo)

	mux := http.NewServeMux()

	v1.New(mux, usecase)

	if err := http.ListenAndServe(cfg.ServerSocket, mux); err != nil {
		log.Fatal(err)
	}
}
