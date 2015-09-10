package main

import (
	"fmt"
	"time"
)

type Bootcamp struct {
	lat float64
	lon float64
	date time.Time
}

func main() {
	element := Bootcamp{lat:34.012836,lon:-118.495338,date:time.Now()}
	fmt.Println(element)
}