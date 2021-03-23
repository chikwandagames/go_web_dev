package main

import (
	"html/template"
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

type user struct {
	UserName string
	First    string
	Last     string
}

var tpl *template.Template
var dbUsers = map[string]user{} //  user ID, user

// Key string, value string
var dbSessions = map[string]string{} //  session ID, user ID
// Alternatively
// var dbSessions1 = make(map[string]string)

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {

	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		// If no cookie present, create one and a uuid
		sID := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
	}

	var u user
	// if the user exists already i.e. we've come back to this page, get user
	// Pull out the userID by session ID
	if userID, ok := dbSessions[c.Value]; ok {
		// Pull out user by userID
		u = dbUsers[userID]
	}

	// process form submission
	// If form is being submitted
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")

		// Assign values to user
		u = user{un, f, l}

		// Put cookie value key (sessionID), and value == username into dbSessions map
		dbSessions[c.Value] = un
		// Put that username Key un and vulue user "u" into dbUsers map
		dbUsers[un] = u
	}

	tpl.ExecuteTemplate(w, "index.html", u)
}

func bar(w http.ResponseWriter, req *http.Request) {

	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	// Get the username
	un, ok := dbSessions[c.Value]
	// If username not found
	if !ok {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	u := dbUsers[un]
	tpl.ExecuteTemplate(w, "bar.html", u)
}

// map examples with the comma, ok idiom
// https://play.golang.org/p/OKGL6phY_x
// https://play.golang.org/p/yORyGUZviV
