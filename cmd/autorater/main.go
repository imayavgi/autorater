package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/imayavgi/autorater/internal/pkg/models"
)

// Values ...
type Values struct {
	Models []Model `json:"values"`
}

// Model ...
type Model struct {
	Name     string   `json:"model"`
	Feedback []string `json:"feedback"`
}

type rating float32

const (
	extraPositive rating = 1.2
	positive      rating = 0.6
	negative      rating = -0.6
	initial       rating = 5.0
	extraNegative rating = -1.2
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
	generateRating()

	// Print ratings for the different vehicles
	for _, veh := range inventory {
		switch v := veh.(type) {
		case models.Car:
			v.CarDetails()
		case models.Bike:
			v.BikeDetails()
		case models.Truck:
			v.TruckDetails()
		default:
			fmt.Printf("Are you sure this Vehicle Type exists")
		}
	}
}

func readJSONFile() Values {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	jsonFile, err := os.Open(path + "/test/data/feedback.json")

	if err != nil {
		log.Fatal(err)
		//log.Fatal("File not found")
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var content Values
	json.Unmarshal(byteValue, &content)

	return content
}

func generateRating() {
	f := readJSONFile()
	for _, v := range f.Models {
		var vehResult models.FeedbackResult
		var vehRating rating

		for _, msg := range v.Feedback {
			if text := strings.Split(msg, ""); len(text) >= 5 {
				vehRating = 5.0
				vehResult.FeedbackTotal++

				for _, word := range text {
					switch s := strings.Trim(strings.ToLower(word), " ,.,!,?,\t,\n,\r"); s {
					case "pleasure", "impressed", "wonderful", "fantastic", "splendid":
						vehRating += extraPositive
					case "help", "helpful", "thanks", "thank you", "happy":
						vehRating += positive
					case "not helpful", "sad", "angry", "improve", "annoy":
						vehRating += negative
					case "pathetic", "bad", "worse", "unfortunately", "agitated", "frustrated":
						vehRating += extraNegative
					}
				}

				switch {
				case vehRating > 8.0:
					vehResult.FeedbackPositive++
				case vehRating >= 4.0 && vehRating <= 8.0:
					vehResult.FeedbackNeutral++
				case vehRating < 4.0:
					vehResult.FeedbackNegative++
				}
			}
		}
		models.VehicleResult[v.Name] = vehResult
	}
}
