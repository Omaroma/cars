package db

import (
	"database/sql"
	"log"
)

func createCarTable(db *sql.DB) {
	createCarTableSQL := `CREATE TABLE car (
		"registration" TEXT NOT NULL PRIMARY KEY,		
		"model" TEXT,
		"mileage" integer,
		"rented" integer		
	  );`

	log.Println("Create car table...")
	statement, err := db.Prepare(createCarTableSQL)
	if err != nil {
		log.Fatal(err)
	}
	_, err = statement.Exec()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("car table created")
}

// We are passing db reference connection from main to our method with other parameters
func insertCar(db *sql.DB, registration string, model string, mileage int, rented bool) {
	log.Println("Inserting car record ...")
	insertCarSQL := `INSERT INTO car(registration, model, mileage, rented) VALUES (?, ?, ?, ?)`
	statement, err := db.Prepare(insertCarSQL) // Prepare statement.
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(registration, model, mileage, rented)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func insertCars(sqliteDatabase *sql.DB) {
	insertCar(sqliteDatabase, "BTS810", "Tesla M3", 100, true)
	insertCar(sqliteDatabase, "BTS811", "Tesla M3", 200, true)
	insertCar(sqliteDatabase, "BTS812", "Tesla M3", 300, true)
	insertCar(sqliteDatabase, "BTS813", "Tesla M3", 400, true)
	insertCar(sqliteDatabase, "BTS814", "Tesla M3", 500, true)
	insertCar(sqliteDatabase, "BTS815", "Tesla M3", 600, false)
	insertCar(sqliteDatabase, "BTS816", "Tesla M3", 700, false)
	insertCar(sqliteDatabase, "BTS817", "Tesla M3", 800, false)
	insertCar(sqliteDatabase, "BTS818", "Tesla M3", 900, false)
	insertCar(sqliteDatabase, "BTS819", "Tesla M3", 1000, false)
}
