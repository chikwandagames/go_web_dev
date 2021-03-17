package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
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
	// Get data out of the form
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(r.Form)
	// fmt.Printf("Form type: %T \n", r.Form)

	data := struct {
		Method      string
		URL         *url.URL
		Submissions map[string][]string
		Header      http.Header
	}{
		r.Method,
		r.URL,
		r.Form, // form data
		r.Header,
	}
	// fmt.Printf("data: %v", data)

	tmpl.ExecuteTemplate(w, "index.gohtml", data)
}
