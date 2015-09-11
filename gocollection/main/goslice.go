package main

import "fmt"

func main() {
	mySlice := []int{1,2,3,4,5,6,7,8,9}
	fmt.Println(mySlice)
	fmt.Println(mySlice[1:4])
	fmt.Println(mySlice[:3])
	fmt.Println(mySlice[4:])
}