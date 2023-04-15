package main

import (
	"context"
	"log"
	"net/http"

	v1 "github.com/v1adhope/web-service-employees/internal/controller/http/v1"
	"github.com/v1adhope/web-service-employees/internal/usecase"
	"github.com/v1adhope/web-service-employees/internal/usecase/repository"
	"github.com/v1adhope/web-service-employees/pkg/mongodb"
)

const (
	_uri     = "mongodb://0.0.0.0:27017/employeeStorage?timeoutMS=10000&maxPoolSize=99"
	_dbName  = "employeeStorage"
	_colName = "employee"
)

func main() {
	mongoClient, err := mongodb.New(context.Background(), _uri)
	if err != nil {
		log.Fatal(err)
	}
	defer mongoClient.Close(context.Background())

	repo := repository.New(mongoClient.GetCollecion(_dbName, _colName))

	usecase := usecase.New(repo)

	mux := http.NewServeMux()

	v1.New(mux, usecase)

	if err := http.ListenAndServe(":8090", mux); err != nil {
		log.Fatal(err)
	}
}
