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

func GetAllVehicles() []Vehicle {
	mu.RLock()
	defer mu.RUnlock()
	storageVar := VehicleStorage
	return storageVar
}

func GetVehicleById(id int) (Vehicle, error) {
	mu.RLock()
	defer mu.RUnlock()

	for _, car := range VehicleStorage {
		if car.Id == id {
			return car, nil
		}
	}
	return Vehicle{}, errors.New("vehicle not found, check Id field")
}

func CreateVehicle(v Vehicle) (Vehicle, error) {
	mu.Lock()
	defer mu.Unlock()

	v.Id = nextId
	nextId++
	VehicleStorage = append(VehicleStorage, v)
	return v, nil
}

func UpdateVehicleById(id int, v Vehicle) (Vehicle, error) {
	mu.Lock()
	defer mu.Unlock()

	for i, car := range VehicleStorage {
		if car.Id == id {
			v.Id = id
			VehicleStorage[i] = v
			return v, nil
		}
	}
	return v, errors.New("Vehicle update failed")

}

func DeleteVehicleById(id int) error {
	mu.Lock()
	defer mu.Unlock()
	for i, car := range VehicleStorage {
		if car.Id == id {
			VehicleStorage = append(VehicleStorage[:i], VehicleStorage[i+1:]...)
			return nil
		}
	}
	return errors.New("No vehicles found for this id")
}
