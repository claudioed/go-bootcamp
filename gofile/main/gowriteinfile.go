package main

import (
	"fmt"
	"os"
)

func main() {
	w, err := os.Create("./resources/long-file.txt")
	if err != nil { fmt.Println("Error on create file!")}
	for index := 0; index <= 100000; index++ {
		indexFormatted := fmt.Sprintf("%012d", index) + "\n"
		line := indexFormatted
		w.WriteString(line)
	}
	defer w.Close()
	fmt.Println("File Created!")
}
