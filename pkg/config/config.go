package config

import (
	"os"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
)

func Load() {
	godotenv.Load()
}

func AssignEnv() map[string]string {
	var env map[string]string
	env, err := godotenv.Read()
	if err != nil {
		log.Info("Can't read .env file, looking from OS env...")
		env = map[string]string{
			"APP_PORT":         os.Getenv("APP_PORT"),
			"MYSQL_URI":        os.Getenv("MYSQL_URI"),
			"MYSQL_TIMEOUT":    os.Getenv("MYSQL_TIMEOUT"),
			"MYSQL_MAX_POOL":   os.Getenv("MYSQL_MAX_POOL"),
			"MYSQL_MIN_POOL":   os.Getenv("MYSQL_MIN_POOL"),
			"MYSQL_DEBUG_MODE": os.Getenv("MYSQL_DEBUG_MODE"),
		}
	}
	return env
}

func ValidateEnv(appConfig *AppConfig) {
	validate := validator.New()
	if err := validate.Struct(appConfig); err != nil {
		panic(err)
	}
}

func GetString(name string, env map[string]string) string {
	return env[name]
}

func GetInt(name string, env map[string]string) (value int) {
	valueStr := env[name]
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	} else {
		panic(err)
	}
}

func GetBool(name string, env map[string]string) (value bool) {
	valueStr := env[name]
	if value, err := strconv.ParseBool(valueStr); err == nil {
		return value
	} else {
		panic(err)
	}
}
