package main

import (
	"fmt"
	"os"
)

func main() {
	w,err := os.Create("users.txt")

	if err != nil{
		fmt.Println("Error on create file!")
	}

	defer w.Close()

	fmt.Println("File Created!")

}