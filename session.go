package main

import (
	"net/http"
)

func getUser(req *http.Request) user {
	var u user

	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		return u
	}

	// if the user exists already, get user
	if un, ok := dbSessions[c.Value]; ok {
		u = dbUsers[un]
	}
	return u
}

func alreadyLoggedIn(req *http.Request) bool {
	c, err := req.Cookie("session")
	// If no cookie, return false,
	if err != nil {
		return false
	}

	// Get usermame by ID
	un := dbSessions[c.Value]
	// If user with ID c.Value (cookie.value) is found the ok == true
	_, ok := dbUsers[un]
	return ok
}
