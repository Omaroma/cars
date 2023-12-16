package db

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

const DB = "sqlite-database.db"

func InitDB() {
	if dbExist() {
		return
	}
	sqliteDatabase := dbConnection()
	createDB()
	createCarTable(sqliteDatabase)
	insertCars(sqliteDatabase)
}

func dbExist() bool {
	if _, err := os.Stat(DB); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}

func dbConnection() *sql.DB {
	sqliteDatabase, err := sql.Open("sqlite3", DB)
	if err != nil {
		log.Fatal(err.Error())
	}
	return sqliteDatabase
}

func createDB() {
	file, err := os.Create("sqlite-database.db") // Create SQLite file
	if err != nil {
		log.Fatal(err.Error())
	}
	err = file.Close()
	if err != nil {
		log.Println("error closing db file", err)
	}
	log.Println("sqlite-database.db created")
}
