package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

type Config struct {
	Port   string
	DBName string
}

func NewConfig() (*Config, error) {
	port, exist := os.LookupEnv("PORT")
	if !exist {
		return nil, fmt.Errorf("PORT not set in .env")
	}
	dbName, exist := os.LookupEnv("DB_NAME")
	if !exist {
		return nil, fmt.Errorf("DB_NAME not set in .env")
	}

	return &Config{
		Port:   port,
		DBName: dbName,
	}, nil
}

func Database() *sql.DB {
	conf, err := NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	db, err := sql.Open("sqlite3", conf.DBName)
	if err != nil {
		log.Fatal(db)
	}

	return db
}
