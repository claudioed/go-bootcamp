package main

import "fmt"

type Car struct {
	model string
	factory string
}

func NewCar(model string,factory string) (car *Car)  {
	fmt.Println("Creating car....")
	if(len(model) == 0 && len(factory) == 0){
		panic("Error")
	}
	car = &Car{model:model,factory:factory}
	return car
}

func (*Car) Run() {
	fmt.Println("Runnn.....")
}

func main() {
	car:= NewCar("Celta","Chevrolet")
	car.Run()
}

