package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", numberOfSiteVisits)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func numberOfSiteVisits(res http.ResponseWriter, req *http.Request) {

	cookie, err := req.Cookie("my-cookie")

	if err == http.ErrNoCookie {
		// If cookie not found, create one
		cookie = &http.Cookie{
			Name:  "my-cookie",
			Value: "0",
			Path:  "/",
		}
	}

	//  strconv.Atoi(), string convert asci to int
	// convert from string to int
	count, err := strconv.Atoi(cookie.Value)
	fmt.Printf("Cookie value type: %T, count type: %T \n", cookie.Value, count)
	if err != nil {
		log.Fatalln(err)
	}

	count++
	// Conver count back to (asci) string, and set that to the cookie value
	cookie.Value = strconv.Itoa(count)

	http.SetCookie(res, cookie)

	io.WriteString(res, "Visited: "+cookie.Value+" times")
}
