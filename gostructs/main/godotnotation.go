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
	event := Bootcamp{lat:31,lon:24}
	event.date = time.Now()
	fmt.Printf("Event on %s, in location (%f, %f)",event.date, event.lat, event.lon)
}