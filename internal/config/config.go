package config

import (
	"os"
	"sync"
)

const (
	_mongoConURL  = "APP_MONGO_CONURL"
	_serverSocket = "APP_SERVER_SOCKET"
)

type Config struct {
	MongoConURL, ServerSocket string
}

var (
	once sync.Once
	cfg  *Config
)

func GetConfig() *Config {
	once.Do(
		func() {
			cfg = &Config{
				MongoConURL:  "mongodb://storage:27017/employeeStorage?timeoutMS=10000",
				ServerSocket: "0.0.0.0:8090",
			}

			fillVarByEnvKey(_mongoConURL, &cfg.MongoConURL)
			fillVarByEnvKey(_serverSocket, &cfg.ServerSocket)
		})

	return cfg
}

func fillVarByEnvKey(key string, placeholder *string) {
	if env := os.Getenv(key); env != "" {
		*placeholder = env
	}
}
