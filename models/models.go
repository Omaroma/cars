package models

type Car struct {
	Model        string `json:"model"`
	Registration string `json:"registration"`
	MileAge      int    `json:"mileage"`
	Rented       bool   `json:"rented"`
}

type RentCar struct {
	DistanceDriven int `json:"DistanceDriven"`
}
