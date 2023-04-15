package main

import (
	"context"
	"log"
	"net/http"
	"os"

	v1 "github.com/v1adhope/web-service-employees/internal/controller/http/v1"
	"github.com/v1adhope/web-service-employees/internal/usecase"
	"github.com/v1adhope/web-service-employees/internal/usecase/repository"
	"github.com/v1adhope/web-service-employees/pkg/mongodb"
)

const (
	_mongoConStr  = "APP_MONGO_CONSTR"
	_serverSocket = "APP_SERVER_SOCKET"

	_dbName  = "employeeStorage"
	_colName = "employee"
)

func main() {
	var conStr, serverSocket string

	getEnv(_mongoConStr, &conStr)
	getEnv(_serverSocket, &serverSocket)

	mongoClient, err := mongodb.New(context.Background(), conStr)
	if err != nil {
		log.Fatal(err)
	}
	defer mongoClient.Close(context.Background())

	repo := repository.New(mongoClient.GetCollecion(_dbName, _colName))

	usecase := usecase.New(repo)

	mux := http.NewServeMux()

	v1.New(mux, usecase)

	if err := http.ListenAndServe(serverSocket, mux); err != nil {
		log.Fatal(err)
	}
}

func getEnv(key string, placeholder *string) {
	if env := os.Getenv(key); env != "" {
		*placeholder = env
	}
}
