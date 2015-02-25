package main

import "net/http"

func main() {
	http.ListenAndServeTLS(":8443", "crt.pem", "key.pem", nil) // HL
}
