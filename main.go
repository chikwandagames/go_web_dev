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

	// Because handle take a http.HandlerFunc type, we can cast d and c
	// to type HandlerFunc, http.HandlerFunc implements
	// ServerHTTP(ResponseWriter, *Request), there for it is a handler
	// Handle() wants a handler

	// the cast is possible because d and c have the same signature as
	// ServerHTTP

	http.Handle("/dog/", http.HandlerFunc(d))
	http.Handle("/cat", http.HandlerFunc(c))

	// If you pass nil, ListenAndServe will use DefaultServeMux
	http.ListenAndServe(":8080", nil)
}

func d(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello Chihuahua")
}

func c(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello Sphynx")
}
