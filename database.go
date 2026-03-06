package main

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	db *pgxpool.Pool
)

func InitDB(str string) error {
	var err error
	db, err = pgxpool.New(context.Background(), str)
	if err != nil {
		return errors.New("failed to init db")
	}

	err = db.Ping(context.Background())
	if err != nil {
		return errors.New("failed to ping db")
	}
	return nil
}

func GetAllVehicles() ([]Vehicle, error) {
	var vehicleSlice []Vehicle
	rows, err := db.Query(
		context.Background(),
		"SELECT id, brand, model, engine, transmission, hp_amount, fuel_type, year_of_prod, mileage, description, main_photo, photos FROM vehicles")
	if err != nil {
		return nil, errors.New("cant make a query")
	}
	defer rows.Close()
	for rows.Next() {
		var v Vehicle
		err = rows.Scan(&v.Id, &v.Brand, &v.Model, &v.Engine, &v.Transmission, &v.HPAmount, &v.FuelType, &v.YearOfProd, &v.Mileage, &v.Description, &v.MainPhoto, &v.Photos)
		if err != nil {
			return nil, errors.New("cant find a row")
		}
		vehicleSlice = append(vehicleSlice, v)
	}
	return vehicleSlice, nil
}

func GetVehicleById(id int) (Vehicle, error) {
	var v Vehicle
	row := db.QueryRow(context.Background(), "SELECT id, brand, model, engine, transmission, hp_amount, fuel_type, year_of_prod, mileage, description, main_photo, photos FROM vehicles WHERE id = $1", id)
	err := row.Scan(&v.Id, &v.Brand, &v.Model, &v.Engine, &v.Transmission, &v.HPAmount, &v.FuelType, &v.YearOfProd, &v.Mileage, &v.Description, &v.MainPhoto, &v.Photos)
	if err != nil {
		return Vehicle{}, errors.New("vehicle not found, check Id field")
	}
	return v, nil
}

func CreateVehicle(v Vehicle) (Vehicle, error) {
	row := db.QueryRow(context.Background(),
		"INSERT INTO vehicles (brand, model, engine, transmission, hp_amount, fuel_type, year_of_prod, mileage, description, main_photo, photos) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id", v.Brand, v.Model, v.Engine, v.Transmission, v.HPAmount, v.FuelType, v.YearOfProd, v.Mileage, v.Description, v.MainPhoto, v.Photos)
	err := row.Scan(&v.Id)
	if err != nil {
		return Vehicle{}, errors.New("can't create vehicle")
	}
	return v, nil
}

func UpdateVehicleById(id int, v Vehicle) (Vehicle, error) {
	cT, err := db.Exec(context.Background(), "UPDATE vehicles SET brand=$1, model=$2, engine=$3, transmission=$4, hp_amount=$5, fuel_type=$6, year_of_prod=$7, mileage=$8, description=$9, main_photo=$10, photos=$11 WHERE id=$12", v.Brand, v.Model, v.Engine, v.Transmission, v.HPAmount, v.FuelType, v.YearOfProd, v.Mileage, v.Description, v.MainPhoto, v.Photos, id)
	if err != nil {
		return Vehicle{}, errors.New("can't update vehicle")
	}
	numOfhanges := cT.RowsAffected()
	if numOfhanges == 0 {
		return Vehicle{}, errors.New("can't find vehicle for this id")
	}
	return v, nil
}

func DeleteVehicleById(id int) error {
	cT, err := db.Exec(context.Background(), "DELETE FROM vehicles WHERE id=$1", id)
	if err != nil {
		return errors.New("can't delete vehicle")
	}
	numOfhanges := cT.RowsAffected()
	if numOfhanges == 0 {
		return errors.New("vehicle not found for id")
	}
	return nil
}
