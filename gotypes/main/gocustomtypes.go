package main

import (
	"fmt"
)

type ListaDeCompras []string

func main() {
	lista := make(ListaDeCompras,4);
	lista[0] = "Alface"
	lista[1] = "Pepino"
	lista[2] = "Arroz"
	lista[3] = "Frango"
	vegetais,carnes,outros := lista.Categorizar();
	fmt.Println(vegetais,carnes,outros)
}

func (lista ListaDeCompras)Categorizar()([]string,[]string,[]string)  {
	var vegetais []string;
	var carnes []string;
	var outros []string;
	for _, e := range lista {
		switch e {
		case "Alface","Pepino":
			vegetais = append(vegetais,e)
		case "Atum","Frango":
			carnes = append(carnes,e)
		default:
			outros = append(outros,e)
		}
	}
	return vegetais,carnes,outros
}