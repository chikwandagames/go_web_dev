package main

import (
	"fmt"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

// To create a session
// We give a user a unique identifier (number),
// Store that in cookie
// Then when the user makes a request to the server
// the will send that unique id
// The server then take that unique id, to determaine
// who is communicating with the server

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session")
	if err != nil {
		id := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session",
			Value: id.String(),
			// Secure: true,
			HttpOnly: true,
			Path:     "/",
		}
		http.SetCookie(w, cookie)
	}
	fmt.Println(cookie)
}
