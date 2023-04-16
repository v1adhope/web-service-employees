package config

import (
	"os"
	"sync"
)

const (
	_mongoConStr  = "APP_MONGO_CONSTR"
	_serverSocket = "APP_SERVER_SOCKET"
)

type Config struct {
	ConStr, ServerSocket string
}

var (
	once sync.Once
	cfg  *Config
)

func GetConfig() *Config {
	once.Do(
		func() {
			cfg = &Config{
				ConStr:       "mongodb://storage:27017/employeeStorage?timeoutMS=10000",
				ServerSocket: "0.0.0.0:8090",
			}

			getEnv(_mongoConStr, &cfg.ConStr)
			getEnv(_serverSocket, &cfg.ServerSocket)
		})

	return cfg
}

func getEnv(key string, placeholder *string) {
	if env := os.Getenv(key); env != "" {
		*placeholder = env
	}
}
