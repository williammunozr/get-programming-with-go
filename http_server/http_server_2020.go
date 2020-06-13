package main

import (
	"fmt"
	"net/http"
)

func welcome_message(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "Welcome to HTTP Server Version 2.0\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	http.HandleFunc("/", welcome_message)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8081", nil)
}
