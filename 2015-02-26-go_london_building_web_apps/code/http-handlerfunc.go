package main

import (
	"log"
	"net/http"
)

func main() {
	hi := func(w http.ResponseWriter, r *http.Request) { // HL
		w.Write([]byte("<h1>Hello, world.</h1>")) // HL
	} // HL

	log.Fatal(http.ListenAndServe(":8080", http.HandlerFunc(hi)))
}
