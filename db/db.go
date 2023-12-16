package db

import (
	"github.com/Omaroma/cars/models"
)

func GetAllCars() (cars []models.Car) {
	sqliteDatabase := dbConnection()
	rows, _ := sqliteDatabase.Query("SELECT * FROM car")
	defer rows.Close()
	for rows.Next() {
		var car models.Car
		rows.Scan(&car.Registration, &car.Model, &car.MileAge, &car.Rented)
		cars = append(cars, car)
	}
	return cars
}

//CarExist check if care exists
func CarExist(registration string) (exist bool) {
	sqliteDatabase := dbConnection()
	sqliteDatabase.QueryRow("SELECT (SELECT 1 FROM car WHERE registration = ?)", registration).Scan(&exist)
	return exist
}

func GetCar(registration string) (car models.Car) {
	sqliteDatabase := dbConnection()
	row := sqliteDatabase.QueryRow("SELECT * FROM car WHERE registration = ?", registration)
	row.Scan(&car.Registration, &car.Model, &car.MileAge, &car.Rented)
	return car
}

func InsertCar(car models.Car) {
	sqliteDatabase := dbConnection()
	sqliteDatabase.Exec("INSERT INTO car VALUES (?, ?, ?, ?)", car.Registration, car.Model, car.MileAge, car.Rented)
}

func RentCar(car models.Car) {
	sqliteDatabase := dbConnection()
	sqliteDatabase.Exec("UPDATE car SET rented = ? WHERE registration = ?", car.Rented, car.Registration)
}

func ReturnCar(car models.Car) {
	sqliteDatabase := dbConnection()
	sqliteDatabase.Exec("UPDATE car SET rented = ?, mileage = ? WHERE registration = ?", car.Rented, car.MileAge, car.Registration)
}
