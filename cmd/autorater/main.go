package main

import (
	"fmt"

	"github.com/imayavgi/autorater/internal/pkg/feedback"
	"github.com/imayavgi/autorater/internal/pkg/models"
)

var inventory []models.Vehicle

func init() {

	inventory = []models.Vehicle{
		models.Bike{Model: "FTR 1200", Make: "Indian"},
		models.Bike{Model: "Iron 1200", Make: "Harley"},
		models.Car{Model: "Sonata", Make: "Hyundai", Type: "Sedan"},
		models.Car{Model: "SantaFe", Make: "Hyundai", Type: "SUV"},
		models.Car{Model: "Civic", Make: "Honda", Type: "Hatchback"},
		models.Car{Model: "A5", Make: "Audi", Type: "Coupe"},
		models.Car{Model: "Mazda6", Make: "Mazda", Type: "Sedan"},
		models.Car{Model: "CRV", Make: "Honda", Type: "SUV"},
		models.Car{Model: "Camry", Make: "Toyota", Type: "Sedan"},
		models.Truck{Model: "F-150", Make: "Ford", Type: "Truck"},
		models.Truck{Model: "RAM1500", Make: "Dodge", Type: "Truck"}}

}

func main() {

	// Generate ratings for the different vehicles
	report := feedback.ProcessRatingFeed()

	// Print ratings for the different vehicles
	for _, veh := range inventory {
		switch v := veh.(type) {
		case models.Car:
			v.CarDetails(report)
		case models.Bike:
			v.BikeDetails(report)
		case models.Truck:
			v.TruckDetails(report)
		default:
			fmt.Printf("Are you sure this Vehicle Type exists")
		}
	}
}
