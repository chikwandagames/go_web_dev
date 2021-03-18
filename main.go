package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {

	// Set hearder Content-Type to text/html
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// WriteString to the ResponseWriter, a sting with a paht to an image
	io.WriteString(w, `
	<!--not serving from our server-->
	<img src="https://upload.wikimedia.org/wikipedia/commons/6/6e/Golde33443.jpg">
	`)
}
