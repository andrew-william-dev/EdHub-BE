package configs

import (
	"database/sql"
	"log"
	"os"
)

func ConnectDB() *sql.DB {
	connStr := os.Getenv("CONN_STR")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect with the DataBase:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping DB:", err)
	}

	return db
}
