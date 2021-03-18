package main

import (
	"fmt"
	"net/http"
)

// http.Handeer

// Anything the implements ServerHTTP(ResponseWriter, *Request)

type hotdog int

func main() {
	var d hotdog
	// ListenAndServe(addr string, handler Handler)
	http.ListenAndServe(":8080", d)

}

// Attaching the ServerHTTP(ResponseWriter, *Request) function to
// type hotdog makes hotdog be of type http.Handler
func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Wisdom", "this is from wisdom")
	// If we change the Content-Type to text/html, then the browser
	// will interpret the data as html
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "<h1>Any code you want in this func </h1>")
}
