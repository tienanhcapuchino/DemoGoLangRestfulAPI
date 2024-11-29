package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type PostgreS struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
	Debug    bool
}

var Envs = initConfig()

func initConfig() PostgreS {
	godotenv.Load()

	return PostgreS{
		Host: getENV("POSTGRES_HOST"),
		Port: getENV("POSTGRES_PORT"),
		User: getENV("POSTGRES_USER"),
		Password: getENV("POSTGRES_PASSWORD"),
		DBName: getENV("POSTGRES_DB"),
	}
}

func ConnectionString(p PostgreS) string {
	return fmt.Sprintf("postgresql://%v:%v@%v:%v/%v", p.User, p.Password, p.Host, p.Port, p.DBName)
}

func getENV(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return ""
}