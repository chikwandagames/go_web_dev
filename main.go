package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/toby.jpg", dogPic)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// src="/toby.jpg" invokes the 	http.HandleFunc("/toby.jpg", dogPic) func
	// this then runs the dogPic function, which serves the image using
	//
	io.WriteString(w, `<img src="/toby.jpg">`)
}

func dogPic(w http.ResponseWriter, req *http.Request) {
	// os.Open give a pointer to the file
	f, err := os.Open("toby.jpg")
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	defer f.Close()

	// Stat() returns the FileInfo structure describing the file
	// If there is an error, it will be of type *PathError
	fi, err := f.Stat()
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}

	// fi.ModTime() is the last time the file was modified
	http.ServeContent(w, req, f.Name(), fi.ModTime(), f)
}
