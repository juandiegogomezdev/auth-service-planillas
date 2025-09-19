package config

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type configDB struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

func LoadDBConfig() configDB {
	return configDB{
		Host:     GetEnv("DB_HOST", "localhost"),
		Port:     GetEnv("DB_PORT", "5433"),
		Username: GetEnv("DB_USER", "juan"),
		Password: GetEnv("DB_PASSWORD", "tunclave"),
		Database: GetEnv("DB_NAME", "juan"),
	}
}

func ConnectDB() *sqlx.DB {
	config := LoadDBConfig()
	host := config.Host
	port := config.Port
	user := config.Username
	password := config.Password
	dbname := config.Database

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	// Verify the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successful connection to PostgreSQL 🚀")
	return db
}
