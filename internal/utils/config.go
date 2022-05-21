package utils

import (
	"log"
	"os"
)

// just enviroment loader
type AppConfig struct {
	DbUser      string
	DbPwd       string
	DbPort      string
	DbAddr      string
	DbSchema    string
	SerivcePort string
}

func LoadCnfFromEnv() *AppConfig {
	return &AppConfig{
		DbUser:      loadEnvByKey("DB_USER"),
		DbPwd:       loadEnvByKey("DB_PWD"),
		DbPort:      loadEnvByKey("DB_PORT"),
		DbAddr:      loadEnvByKey("DB_ADDR"),
		DbSchema:    loadEnvByKey("DB_SCHEMA"),
		SerivcePort: loadEnvByKey("SERVICE_PORT"),
	}
}

func loadEnvByKey(key string) (value string) {
	val, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("cannot load env value by key = %s", key)
	}
	return val
}
