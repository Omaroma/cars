package main

import (
	"github.com/Omaroma/cars/db"
	"github.com/Omaroma/cars/middleware"
	"log"
	"net/http"
)

func main() {
	db.InitDB()

	router := middleware.GetRouter()
	port := middleware.Port
	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Fatal("can't start server", err)
	}

	log.Println("server started and listing to port : ", port)
}
