package main

import "fmt"

func main() {
	c:= make(chan int,3)
	go produzirCorrect(c)
	for valor := range c {
		fmt.Println(valor)
	}
}

func produzirCorrect(c chan int) {
	c <- 1
	c <- 2
	c <- 3
	close(c)
}
