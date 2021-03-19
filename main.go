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
	// Retreev
	v := req.FormValue("first_name")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// When we use POST, form data is sent via the request body
	// When you hit submit, the page refresh renders the page with first_name v
	io.WriteString(w, `
	<form method="POST">
	 <input type="text" name="first_name" autocomplete="off" placeholder="first name">
	 <input type="submit">
	</form>
	<br> <p> your name is: `+v)
}
