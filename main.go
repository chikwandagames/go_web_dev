package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("first_name")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// On submit the form data will be passed to server via the URL as query strings
	io.WriteString(w, `
	<form method="GET">
	<input type="text" name="first_name" autocomplete="off" placeholder="first name">
	<input type="submit">
	</form>
	<br> <p> your name is: `+v)
}
