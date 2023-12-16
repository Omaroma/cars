package services

import (
	"encoding/json"
	"fmt"
	"github.com/Omaroma/cars/db"
	"github.com/Omaroma/cars/models"
	"github.com/gorilla/mux"
	"net/http"
)

// ListCars returns a list of all the cars registered.
//Te response should include every car's model, its registration number, its mileage and its rental status are displayed.
func ListCars(w http.ResponseWriter, r *http.Request) {
	cars := db.GetAllCars()
	data, _ := json.Marshal(cars)
	w.Write(data)
}

//AddCar The API takes the new car to register, if the car already exists (the registration number already exists) the API reports an error
//, if not the API stores the new car and returns its inserted id
func AddCar(w http.ResponseWriter, r *http.Request) {
	var car models.Car
	err := json.NewDecoder(r.Body).Decode(&car)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if db.CarExist(car.Registration) {
		http.Error(w, "car already exists.", http.StatusBadRequest)
		return
	}

	db.InsertCar(car)
	w.Write([]byte(fmt.Sprintf("%v", car.Registration)))
}

//RentCar The API uses the registration number of the car to be rented.
//If the car does not exist, the API reports an error; if it is already rented,
//the API indicates that it is already rented otherwise the car is marked as being rented.
func RentCar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var car models.Car
	if !db.CarExist(vars["registration"]) {
		http.Error(w, "car doesn't exist.", http.StatusNotFound)
		return
	}

	car = db.GetCar(vars["registration"])
	if car.Rented {
		http.Error(w, "car already rented.", http.StatusBadRequest)
		return
	}

	car.Rented = true
	db.RentCar(car)
	w.Write([]byte("car rented successfully."))
}

//ReturnCar The API uses the registration number of the car to be returned.
//If the car does not exist, the program reports an error; if the car was not marked as being rented,
//the program stipulates it otherwise the program takes the number of kilometers driven and adds them to the mileage of the car.
//The car is then marked as available
func ReturnCar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var car models.Car
	if !db.CarExist(vars["registration"]) {
		http.Error(w, "car doesn't exist.", http.StatusNotFound)
		return
	}

	car = db.GetCar(vars["registration"])
	if !car.Rented {
		http.Error(w, "car wasn't rented.", http.StatusBadRequest)
		return
	}

	var RentedCar models.RentCar
	err := json.NewDecoder(r.Body).Decode(&RentedCar)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	car.MileAge += RentedCar.DistanceDriven
	car.Rented = false
	db.ReturnCar(car)
	w.Write([]byte("car returned successfully"))
}
