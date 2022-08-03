package config

type AppConfig struct {
	AppPort          string `validate:"required"`
	MySqlUri         string `validate:"required"`
	MySqlTimeout     int    `validate:"required"`
	MySqlMaxPool     int    `validate:"required"`
	MySqlMinPoolSize int    `validate:"required"`
	MySqlDebugMode   bool   `validate:"required"`
}
