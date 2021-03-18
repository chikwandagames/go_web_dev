package main

import (
	"io"
	"net/http"
)

// http.Handeer

// Anything the implements ServerHTTP(ResponseWriter, *Request)
// is of type handler
/*
	type Handler interface {
		ServerHTTP(ResponseWriter, *Request)
	}
*/

func main() {

	// Because handle func take a function with the signature
	// myFunc(w http.ResponseWriter, r *http.Request)
	http.HandleFunc("/dog/", d)
	http.HandleFunc("/cat", c)

	// If you pass nil, ListenAndServe will use DefaultServeMux
	http.ListenAndServe(":8080", nil)
}

func d(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello Chihuahua")
}

func c(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello Sphynx")
}
