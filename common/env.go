package common

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func ParseEnv(key string, required bool, dft string) string {
	_ = godotenv.Load()
	value := os.Getenv(key)
	if value == "" && required {
		panic(fmt.Sprintf("Environment variable not found: %v", key))
	} else if value == "" {
		return dft
	}
	return value
}

type Environment struct {
	ApplicationHost string
	Port            string
	Debug           bool
	RedisAddr       string
	RedisPassword   string
	Postgres        string
}

func GetEnvironment() *Environment {
	return &Environment{
		ApplicationHost: ParseEnv("APPLICATION_HOST", false, "0.0.0.0"),
		Port:            ParseEnv("PORT", false, "8000"),
		Debug:           ParseEnv("DEBUG", false, "false") == "true",
		Postgres:        ParseEnv("DATABASE_URL", true, "false"),
		RedisAddr:       ParseEnv("REDIS", false, "false"),
	}
}
