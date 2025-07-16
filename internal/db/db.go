package db

import (
	"database/sql"
	"os"
	"path/filepath"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "host=79.174.88.248 user=ozonprac password=OzonPrac2025! port=18477 dbname=oprac sslmode=require")
	if err != nil {
		return nil, err
	}
	DB = db

	if err := db.Ping(); err != nil {
		return nil, err
	}

	path, err := filepath.Abs("internal/db/schema.sql")
	if err != nil {
		return nil, err
	}
	schema, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(string(schema))

	return db, err
}

func GetDB() *sql.DB {
	return DB
}

func ReadQueryFromFile(path string) (string, error) {
	pathA, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	query, err := os.ReadFile(pathA)
	if err != nil {
		return "", err
	}

	return string(query), nil
}