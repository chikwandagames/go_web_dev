package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type person struct {
	FirstName  string
	LastName   string
	Subscribed bool
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {

	fName := req.FormValue("first")
	lName := req.FormValue("last")
	// if the checkbox is ticked, then req.FormValue("subscribe") == "on"
	isSub := req.FormValue("subscribe") == "on"

	p1 := person{FirstName: fName, LastName: lName, Subscribed: isSub}

	err := tpl.ExecuteTemplate(w, "index.html", p1)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}
}
