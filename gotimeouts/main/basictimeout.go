package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan bool,1)
	go executar(c)
	fmt.Println("Esperando...")
}

func executar(c chan <- bool)  {
	time.Sleep(5* time.Second)
	c <- true
}