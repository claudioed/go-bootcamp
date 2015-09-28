package main


import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter,req * http.Request)  {
	fmt.Fprintf(w,"My First HTTP Req.")
}

func main() {
	http.HandleFunc("/",handler)
	http.ListenAndServe(":8080",nil)
}