package main

import (
	"io"
	"net/http"
)

func main() {
	// FileServer takes a directory, current dir "." in this case
	// and returns a handler
	// http.FileServer(http.Dir(".")) means serve everything the
	// current directory
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/dog/", dog)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// When the browser runs the html <img src="/toby.jpg">
	// it look for /toby.jpg, the closest match is Handler("/")
	// Then http.FileServer checks for the /toby.jpg file finds
	// it and serves it

	io.WriteString(w, `<img src="/toby.jpg">`)
}
