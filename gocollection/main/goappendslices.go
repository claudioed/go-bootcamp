package main

import "fmt"

func main() {
	cities := []string{"Sao Paulo","Los Angeles"}
	otherCities := []string{"Campinas","Americana"}
	cities = append(cities,otherCities...)
	fmt.Printf("%q",cities)
}
