package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/barred", barred)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Print("Your request method at foo: ", req.Method, "\n\n")
}

// The form in barred (index.html) is submitted to bar handler
// <form method="POST" action="/bar">

func bar(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at bar:", req.Method)
	// process form submission here

	// If you hit bar once for the first time, the browser will not
	// allow you to hit bar again,
	// each time you hit bar the browser will automatically redirect you
	// to foo "/" because of http.StatusMovedPermanently.
	// You would need to clear the cache to access bar again
	http.Redirect(w, req, "/", http.StatusMovedPermanently)
}

func barred(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at barred:", req.Method)
	tpl.ExecuteTemplate(w, "index.html", nil)
}
