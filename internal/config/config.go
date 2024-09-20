package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

type Config struct {
	// db configs
	Db DBConfig
}

var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		err := godotenv.Load("internal/config/config.env") // почему-то с точкой не работет?
		if err != nil {
			log.Println("error open config file")
		}
		log.Println("cfg load")
	})
	return &Config{Db: DBConfig{
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("PORT"),
		User:     os.Getenv("USER"),
		Password: os.Getenv("PASSWORD"),
		DbName:   os.Getenv("DBNAME"),
	}}
}
