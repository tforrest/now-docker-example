package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", hanlder)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func hanlder(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello! Your IP Address is %s!", r.RemoteAddr)
}
