package main

import (
	"fmt"
	"net/http"
	"time"
	"log"
)

func indexHandler(w http.ResponseWriter,req *http.Request) {
	t1 := time.Now();
	fmt.Fprintf(w,"Log Http Request Time!!!")
	t2 := time.Now()
	log.Printf("[%s] %q %v\n", req.Method, req.URL.String(), t2.Sub(t1))
}

func main() {
	http.HandleFunc("/",indexHandler)
	http.ListenAndServe(":8080",nil)
}

