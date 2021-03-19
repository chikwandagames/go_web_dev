package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)

	// http://localhost:8080/resources/ will list all files in the assets folder
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//

	// in the <img> tag /resources/ referes to the /resources route
	// toby.jpg is the file in the assets folder
	// http.StripPrefix("/resources") will remover "/resources" from
	// src="/resources/toby.jpg" and leave just /toby.jpg
	// http.FileServer() will look for anything in the ./assets folder
	// in the current dir in this case /toby.jpg

	// So technically ./assets will be concatenated to /toby by http.FileServer
	// After http.StripPrefix has removed /resources from src="/resources/toby.jpg"
	io.WriteString(w, `<img src="/resources/toby.jpg">`)
	io.WriteString(w, `<img src="/resources/auba.jpg">`)
}
