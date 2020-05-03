package models

import (
	"fmt"

	"github.com/imayavgi/autorater/internal/pkg/feedback"
)

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
func (c *Car) CarDetails(report feedback.Report) {
	fmt.Printf("\n%-5v: %-8v: %-12v ", "Car", c.Make, c.Model)
	showRating(c.Model, report)
}

//BikeDetails ...
func (b *Bike) BikeDetails(report feedback.Report) {
	fmt.Printf("\n%-5v: %-8v: %-12v ", "Bike", b.Make, b.Model)
	showRating(b.Model, report)
}

//TruckDetails ...
func (t *Truck) TruckDetails(report feedback.Report) {
	fmt.Printf("\n%-5v: %-8v: %-12v ", "Truck", t.Make, t.Model)
	showRating(t.Model, report)
}

func showRating(model string, report feedback.Report) {
	ratingFound := false

	for m, r := range report {
		if m == model {
			fmt.Printf("Total Ratings:%v\tPositive:%v\tNegative:%v\tNeutral:%v", r.FeedbackTotal, r.FeedbackPositive, r.FeedbackNegative, r.FeedbackNeutral)
			ratingFound = true
		}
	}

	if !ratingFound {
		fmt.Printf("No rating for this vehicle")
	}
}
