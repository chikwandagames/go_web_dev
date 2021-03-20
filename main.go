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

	// The difference from using http.StatusSeeOthe to http.StatusTemporaryRedirect
	// is that the method does not change to GET it stays as POST
	http.Redirect(w, req, "/", http.StatusTemporaryRedirect)
}

func barred(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at barred:", req.Method)
	tpl.ExecuteTemplate(w, "index.html", nil)
}
