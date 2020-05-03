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
	report.ShowRating(c.Model)
}

//BikeDetails ...
func (b *Bike) BikeDetails(report feedback.Report) {
	fmt.Printf("\n%-5v: %-8v: %-12v ", "Bike", b.Make, b.Model)
	report.ShowRating(b.Model)
}

//TruckDetails ...
func (t *Truck) TruckDetails(report feedback.Report) {
	fmt.Printf("\n%-5v: %-8v: %-12v ", "Truck", t.Make, t.Model)
	report.ShowRating(t.Model)
}
