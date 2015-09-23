package main
import (
		"fmt"
		"math"
	)

func main() {
	a := []int{2,-1,0,4}
	b := []int{3,0,-2,5}
	var sub int
	var pot float64
	soma := float64(0)
	var raiz float64

	for idx := 0; idx < len(a); idx++ {
		sub = a[idx] - b[idx]
		pot = math.Pow(float64(sub),2)
		soma = soma + pot
	}

	raiz = math.Sqrt(soma)

	fmt.Println(raiz)

}
