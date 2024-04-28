package config

import (
	"fmt"
	"os"

	"github.com/lipesalin/ecom/env"

)

type Config struct {
	PublicHost     string
	Port           string
	DBUser         string
	DBPassword     string
	DBAddress      string
	DBName         string
	DBHost         string
	DBPort         string
	PsqlConnection string
}

var Envs = initConfig()

func initConfig() Config {
	env.LoadEnv()

	DBHost := getEnv("DB_HOST", "localhost")
	DBPort := getEnv("DB_PORT", "5432")
	DBUser := getEnv("DB_USER", "root")
	DBPassword := getEnv("DB_PASSWORD", "mypassword")

	DBAddress := fmt.Sprintf("%s:%s", DBHost, DBPort)

	psqlconn :=
		fmt.Sprintf(`
			host=%s 
			port=%s 
			user=%s 
			password=%s 
			dbname=%s 
			sslmode=disable`,
			DBHost,
			DBPort,
			DBUser,
			DBPassword,
			"ecom",
		)

	return Config{
		PublicHost:     getEnv("PUBLIC_HOST", "http://localhost"),
		Port:           getEnv("PORT", "8080"),
		DBUser:         DBUser,
		DBPassword:     DBPassword,
		DBAddress:      DBAddress,
		DBName:         getEnv("DB_NAME", "ecom"),
		DBHost:         DBHost,
		DBPort:         DBPort,
		PsqlConnection: psqlconn,
	}
}

func getEnv(key string, fallback string) string {
	if value, find := os.LookupEnv(key); find {
		return value
	}

	return fallback
}
