package configs

import (
	"database/sql"
	"log"
	"os"
)

func ConnectDB() *sql.DB {
	log.Println("Connecting to DB:", os.Getenv("CONN_STR"))
	connStr := "postgresql://neondb_owner:npg_7YKfMJCqWe2l@ep-raspy-violet-a95zz6n1-pooler.gwc.azure.neon.tech/neondb?sslmode=require"
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
