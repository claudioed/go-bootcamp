package main

import "fmt"

type Point struct  {
	x int
	y int
}

var(
	p = Point{1,2}
	q = &Point{1,2}
	r = Point{x:1}
	s = Point{}
)

func main() {
	fmt.Println(p,q,r,s)
}