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
	// Executes a template at, index.html
	http.HandleFunc("/barred", barred)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Print("Your request method at foo: ", req.Method, "\n\n")
}

func bar(w http.ResponseWriter, req *http.Request) {
	// The form in barred (index.html) is submitted to bar handler
	// <form method="POST" action="/bar">
	fmt.Println("Your request method at bar:", req.Method)

	// process form submission here

	// .Set(), will set an new location, which will redirect to "/" foo handler
	w.Header().Set("Location", "/")
	// Becauser we are using http.StatusSeeOther 303 status code, the method
	// will be changed to GET
	w.WriteHeader(http.StatusSeeOther)
}

func barred(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at barred:", req.Method)
	tpl.ExecuteTemplate(w, "index.html", nil)
}
