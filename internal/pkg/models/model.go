package models

import "fmt"

//Vehicle ...
type Vehicle interface {
}

//Car ...
type Car struct {
	Model string
	Make  string
	Type  string
}

//Truck ...
type Truck struct {
	Model string
	Make  string
	Type  string
}

//Bike ...
type Bike struct {
	Model string
	Make  string
}

//CarDetails ...
func (c *Car) CarDetails() {
	fmt.Printf("\n%-5v: %-8v: %-12v ", "Car", c.Make, c.Model)
	showRating(c.Model)
}

//BikeDetails ...
func (b *Bike) BikeDetails() {
	fmt.Printf("\n%-5v: %-8v: %-12v ", "Bike", b.Make, b.Model)
	showRating(b.Model)
}

//TruckDetails ...
func (t *Truck) TruckDetails() {
	fmt.Printf("\n%-5v: %-8v: %-12v ", "Truck", t.Make, t.Model)
	showRating(t.Model)
}

// FeedbackResult ...
type FeedbackResult struct {
	FeedbackTotal    int
	FeedbackPositive int
	FeedbackNegative int
	FeedbackNeutral  int
}

// VehicleResult ...
var VehicleResult map[string]FeedbackResult = make(map[string]FeedbackResult)

func showRating(model string) {
	ratingFound := false

	for m, r := range VehicleResult {
		if m == model {
			fmt.Printf("Total Ratings:%v\tPositive:%v\tNegative:%v\tNeutral:%v", r.FeedbackTotal, r.FeedbackPositive, r.FeedbackNegative, r.FeedbackNeutral)
			ratingFound = true
		}
	}

	if !ratingFound {
		fmt.Printf("No rating for this vehicle")
	}
}
