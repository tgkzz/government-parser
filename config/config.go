package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"parser/internal/models"
)

type Config struct {
	Url models.URL
	Db  DB
}

type DB struct {
	DriverName     string `env:"DRIVERNAME"`
	DataSourceName string
}

func LoadConfig(path string) (Config, error) {
	if err := godotenv.Load(path); err != nil {
		return Config{}, err
	}

	cfg := Config{
		Url: models.URL{
			LotURL:      os.Getenv("LOTURL"),
			DocumentULR: os.Getenv("DOCURL"),
		},
		Db: DB{
			DriverName: os.Getenv("DRIVERNAME"),
			DataSourceName: fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable",
				os.Getenv("DBDRIVER"), os.Getenv("DBUSER"),
				os.Getenv("DBPASS"), os.Getenv("DBHOST"),
				os.Getenv("DBPORT"), os.Getenv("DBNAME"),
			),
		},
	}

	return cfg, nil
}
