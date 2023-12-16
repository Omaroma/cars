package middleware

import (
	"github.com/Omaroma/cars/services"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const Port = ":8080"

var Router = mux.NewRouter()

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func GetRouter() *mux.Router {
	Router.HandleFunc("/cars", services.ListCars).Methods("GET")
	Router.HandleFunc("/cars", services.AddCar).Methods("POST")
	Router.HandleFunc("/cars/{registration}/rentals", services.RentCar).Methods("POST")
	Router.HandleFunc("/cars/{registration}/returns", services.ReturnCar).Methods("POST")
	Router.Use(loggingMiddleware)
	return Router
}
