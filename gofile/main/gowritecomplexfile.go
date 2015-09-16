package main

import (
	"fmt"
	"os"
)

func main() {
	w, err := os.Create("./resources/complex-file.txt")
	if err != nil { fmt.Println("Error on create file!")}
	for index := 0; index <= 100000; index++ {
		user := fmt.Sprintf("user%d",index)
		document := fmt.Sprintf("%011d", index)
		email := fmt.Sprintf("user%d@gmail.com", index)
		indexFormatted := fmt.Sprintf("%s,%s,%s,%012d,(19)-%012d",user,document,email ,index,index) + "\n"
		line := indexFormatted
		w.WriteString(line)
	}
	defer w.Close()
	fmt.Println("File Created!")
}
