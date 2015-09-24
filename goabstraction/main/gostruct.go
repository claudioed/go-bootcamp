package main

import "fmt"

type Arquivo struct {
	nome string
	tamanho float64
	caracteres int
	palavas int
	linhas int
}

func main() {
	ponteiroArquivo := &Arquivo{tamanho:729,nome:"arquivo.txt"}
	fmt.Println(" Antes : ",ponteiroArquivo.nome,ponteiroArquivo.tamanho)
	ponteiroArquivo.tamanho = 20.00
	fmt.Println(" Depois : ",ponteiroArquivo.nome,ponteiroArquivo.tamanho)
}