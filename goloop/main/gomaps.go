package main

import "fmt"

func main() {
	cities := map[string]int{
		"New York": 8888181,
		"Los Angeles":2122121,
		"Chicago":112233,
	}
	for key,value := range cities {
		fmt.Printf("%s has %d inhabitants \n",key,value)
	}
}
