package models

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v3"
)

type Config struct {
	DB_USERNAME string `yaml:"DB_USERNAME"`
	DB_PASSWORD string `yaml:"DB_PASSWORD"`
	DB_HOST     string `yaml:"DB_HOST"`
	DB_PORT     string `yaml:"DB_PORT"`
	DB_NAME     string `yaml:"DB_NAME"`
}

func dsn(config Config) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DB_USERNAME, config.DB_PASSWORD, config.DB_HOST, config.DB_PORT, config.DB_NAME)
}

func Connection() (*sql.DB, error) {
	configFile, err := os.Open("config.yaml")
	if err != nil {
		log.Fatalf("failed to open config file: %v", err)
	}
	defer configFile.Close()

	var config Config
	decoder := yaml.NewDecoder(configFile)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalf("failed to decode config: %v", err)
	}

	// Open database connection
	db, err := sql.Open("mysql", dsn(config))
	if err != nil {
		log.Printf("Error: %s when opening DB", err)
		return nil, err
	}

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(100)
	db.SetConnMaxLifetime(time.Minute * 5)

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	err = db.PingContext(ctx)
	if err != nil {
		log.Printf("Errors %s pinging DB", err)
		return nil, err
	}
	return db, nil
}
