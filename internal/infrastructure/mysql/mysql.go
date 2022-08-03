package mysql

import (
	"fmt"
	"time"

	"github.com/agungwicaksana/privy-pretest/pkg/config"
	"github.com/jmoiron/sqlx"
)

func NewMySQL(dbConfig config.AppConfig) (db *sqlx.DB) {
	db, err := sqlx.Connect("mysql", dbConfig.MySqlUri)
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(dbConfig.MySqlMaxPool)
	db.SetMaxIdleConns(dbConfig.MySqlMinPoolSize)
	db.SetConnMaxLifetime(time.Duration(dbConfig.MySqlTimeout) * time.Second)
	fmt.Println("Database connected")

	return db
}
