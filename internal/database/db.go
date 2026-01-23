package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	connStr := os.Getenv("DATABASE_URL")

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		panic("could not connect to database: " + err.Error())
	}

	if err = DB.Ping(); err != nil {
		panic("could not ping database: " + err.Error())
	}

	fmt.Println("Database connected successfully!")

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTable()
	fmt.Println("Table created!")
}

func Close() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}

func createTable() {
	createPropertiesTable := `
	CREATE TABLE IF NOT EXISTS properties (
		id TEXT PRIMARY KEY,
		title TEXT NOT NULL,
		description TEXT NOT NULL,
		type TEXT NOT NULL,
		status TEXT NOT NULL,
		price BIGINT NOT NULL,
		bedrooms INT NOT NULL,
		bathrooms INT NOT NULL,
		size_sqm INT NOT NULL,
		address TEXT NOT NULL,
		images TEXT[],
		created_at TIMESTAMP NOT NULL DEFAULT NOW()
	)`

	if _, err := DB.Exec(createPropertiesTable); err != nil {
		panic("could not create properties table: " + err.Error())
	}
}
