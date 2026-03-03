package main

type Vehicle struct {
	Id           int      `json:"id"`
	Brand        string   `json:"brand" binding:"required"`
	Model        string   `json:"model" binding:"required"`
	Engine       string   `json:"engine" binding:"required"`
	Transmission string   `json:"transmission" binding:"required"`
	HPAmount     int      `json:"hp_amount" binding:"required"`
	FuelType     string   `json:"fuel_type" binding:"required"`
	YearOfProd   int      `json:"year_of_prod" binding:"required"`
	Mileage      int      `json:"mileage" binding:"required"`
	Description  string   `json:"description" binding:"required"`
	MainPhoto    string   `json:"main_photo"`
	Photos       []string `json:"photos"`
}
