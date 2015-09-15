package main
import "fmt"

func produzirFlow(c chan <- int) {
	c <- 1
	c <- 2
	c <- 3
	close(c)
}

func main() {
	c:= make(chan int,3)
	go produzirFlow(c)
	consumirFlow(c)
}

func consumirFlow(c <- chan int)  {
	for valor := range c{
		fmt.Println(valor)
	}
}

