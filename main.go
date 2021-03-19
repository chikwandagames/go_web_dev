package main

import (
	"fmt"
	"net/http"
)

// visit link
// localahost:8080/?key=whatever

func main() {
	http.HandleFunc("/", foo)
	// http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {

	// FormValue returns the firest value for the name component of the query
	// POST and PUT body params take precedence over URL query string values.
	// FormValues calls ParseMultipartForm and ParseForm if necessary and ignores
	// errors returned by these functions
	// If key is NOT present, FormValues returns an empty string.
	// To access multiple values of the same key, call ParseForm and then
	// inspect Request.Form directly

	// key is from the query i.e. localahost:8080/?key=whatever
	// Here we use FormValue() to retrieve the value from the URL
	v := r.FormValue("key")
	// Set the content-type to html so we can use the <em> tag
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "Your query string value is: <em> %v </em>", v)
}
