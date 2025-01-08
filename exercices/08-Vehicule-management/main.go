package main

import "fmt"

type Vehicule struct {
	Make  string
	Model string
	Year  string
}
type Insurable interface {
	CalculateInsurance() float64
}

type Printable interface {
	Details()
}

type Car struct {
	Vehicule
	NumberOfDoors int
}

func (c Car) CalculateInsurance() float64 {
	return float64(c.NumberOfDoors) * 100
}
func (c Car) Details() {
	fmt.Println(c.Make+" "+c.Model+" "+c.Year, c.NumberOfDoors)
}

type Truck struct {
	Vehicule
	PayloadCapacity float64
}

func (c Truck) CalculateInsurance() float64 {
	return float64(c.PayloadCapacity) * 150
}
func (c Truck) Details() {
	fmt.Println(c.Make+c.Model+c.Year, c.PayloadCapacity)
}

func PrintAll(vehicule []Printable) {
	for _, printable := range vehicule {
		printable.Details()
	}

}

func main() {
	car := Car{
		Vehicule:      Vehicule{Make: "BMW", Model: "Q8", Year: "2025"},
		NumberOfDoors: 10}

	t := Truck{
		Vehicule:        Vehicule{Make: "MAN", Model: "V1", Year: "2025"},
		PayloadCapacity: 2200}

	sliceofVehicule := []Printable{t, car}

	PrintAll(sliceofVehicule)

}
