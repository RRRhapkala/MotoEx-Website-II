package main

import (
	"errors"
	"sync"
)

var (
	VehicleStorage []Vehicle

	nextId int = 1

	mu sync.RWMutex
)

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
