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

type hotdog int
type hotcat int

func main() {
	var d hotdog
	var c hotcat

	// NewServeMux allocates and returns a new ServeMux
	// func NewServeMux() *ServeMux { return new(ServeMux)}
	// Therefore mux is of type *ServeMux
	mux := http.NewServeMux()
	// The trailing / at the end of the path makes a difference
	// "/dog/" will catch  /dog/something/another
	// "/cat" will NOT catch /cat/something
	mux.Handle("/dog/", d)
	mux.Handle("/cat", c)
	http.ListenAndServe(":8080", mux)
}

// Attaching the ServerHTTP(ResponseWriter, *Request) function to
// type hotdog makes hotdog be of type http.Handler
func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello Chihuahua")
}

func (m hotcat) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello Sphynx")
}
