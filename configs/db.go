package configs

import (
	"database/sql"
	"log"
	"os"
)

func ConnectDB() *sql.DB {
	connStr := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect with the DataBase:", err)
	}

	return db
}
