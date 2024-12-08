package configs

import (
	"fmt"
	"os"
	"path"

	"github.com/joho/godotenv"
)

type GetConfigError struct {
	Key string
}

func (e *GetConfigError) Error() string {
	return fmt.Sprintf("cannot find environment variable: %s", e.Key)
}

type Config struct {
	MySQL MySQLConfig
	JWT   JWTConfig
}

func loadEnv() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	filepath := path.Join(dir, "deployments/.env")

	if err := godotenv.Load(filepath); err != nil {
		panic(err)
	}
}

func getEnv(key string) (string, error) {
	value, ok := os.LookupEnv(key)
	if !ok {
		return "", &GetConfigError{
			Key: key,
		}
	}

	return value, nil
}

func NewConfigs() (*Config, error) {
	loadEnv()

	return &Config{
		MySQL: newMySQLConfig(),
		JWT:   newJWTConfig(),
	}, nil
}
