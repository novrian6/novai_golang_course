package main

import "fmt"

// Base Car struct
type Car struct {
	brand string
	year  int
}

// Method to get ElectricCar info
func (e *Car) GetInfo() string {
	return e.brand
}

// ElectricCar struct that embeds Car
type ElectricCar struct {
	Car     // Embed Car
	battery int
}

// Constructor for ElectricCar
func NewElectricCar(brand string, year int, battery int) *ElectricCar {
	return &ElectricCar{
		Car:     Car{brand: brand, year: year},
		battery: battery,
	}
}

// Method to get ElectricCar info
func (e *ElectricCar) GetElectricInfo() string {
	return fmt.Sprintf("%s with battery capacity of %d%%", e.Car.GetInfo(), e.battery)
}

func main() {
	electricCar := NewElectricCar("Tesla", 2023, 85)
	fmt.Println(electricCar.GetElectricInfo())
}
