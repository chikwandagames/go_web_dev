package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

// When you use POST data is sent via the body
// when you use GET the data is sent through the url

type hotdog int

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var d hotdog
	// ListenAndServe(addr string, handler Handler)
	http.ListenAndServe(":8080", d)

}

// Attaching the ServerHTTP(ResponseWriter, *Request) function to
// type hotdog makes hotdog be of type http.Handler
func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r.Form)

	tmpl.ExecuteTemplate(w, "index.gohtml", r.Form)
}
