package config

import (
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func Load() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}

func AssignEnv() map[string]string {
	env, err := godotenv.Read()
	if err != nil {
		panic("Error reading .env file")
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
