package main

import (
	"crypto/sha1"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	uuid "github.com/satori/go.uuid"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	c := getCookie(w, req)

	// process form submission
	if req.Method == http.MethodPost {
		// Get the file
		mf, fh, err := req.FormFile("myFile")
		if err != nil {
			fmt.Println(err)
		}
		defer mf.Close()
		// Get the file extension
		ext := strings.Split(fh.Filename, ".")[1]
		// create sha for file name
		h := sha1.New()
		io.Copy(h, mf)
		fname := fmt.Sprintf("%x", h.Sum(nil)) + "." + ext
		// create new file
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}
		path := filepath.Join(wd, "public", "images", fname)
		nf, err := os.Create(path)
		if err != nil {
			fmt.Println(err)
		}
		defer nf.Close()
		// copy
		mf.Seek(0, 0)
		io.Copy(nf, mf)
		// add filename to this user's cookie
		c = addFileNameValue(w, c, fname)
	}

	// Here we splite the file names by the | delimeter,
	// returning a slice of strings
	xs := strings.Split(c.Value, "|")

	tpl.ExecuteTemplate(w, "index.html", xs)
}

func getCookie(w http.ResponseWriter, req *http.Request) *http.Cookie {
	// Get cookie
	c, err := req.Cookie("session")
	// If no cookie
	if err != nil {
		sID := uuid.NewV4()
		// Create a cookie
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
	}
	return c
}

// Returns a cookie with a string of file names seperated by a |
// added to the cookie value
func addFileNameValue(w http.ResponseWriter, c *http.Cookie, fname string) *http.Cookie {
	// values
	s := c.Value
	if !strings.Contains(s, fname) {
		s += "|" + fname
	}
	c.Value = s
	http.SetCookie(w, c)
	return c
}
