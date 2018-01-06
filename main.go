package main

import (
	"fmt"
	"log"
	"net/http"
)

var visitors int64

func main() {
	http.HandleFunc("/", hanlder)
	log.Println("Starting Server")
	log.Fatal(http.ListenAndServe(":443", nil))
}

func hanlder(w http.ResponseWriter, r *http.Request) {
	visitors = visitors + 1
	fmt.Fprintf(w, "Hello, from Now! %v requests since last deployment", visitors)
}
