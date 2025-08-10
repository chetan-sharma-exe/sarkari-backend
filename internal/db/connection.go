package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func ConnectDatabase(DBurl string) (*sql.DB, error) {
	dsn := DBurl
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("open db: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ping db: %w", err)
	}
	fmt.Println("Database connected!")

	return db, nil
}
