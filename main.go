package main

import (
	"crypto/sha1"
	"fmt"
	"html/template"
	"io"
	"log"
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
	// add route to serve pictures
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	c := getCookie(w, req)

	// process form submission
	if req.Method == http.MethodPost {
		// Get the file, request form file
		// mf multipart.File, fh file.Header
		mf, fh, err := req.FormFile("myFile")
		if err != nil {
			fmt.Println(err)
			log.Fatalln(err)
		}
		// Close multipart file
		defer mf.Close()
		// Get the file extension, by spliting the FileHeader.Filename
		// remove everthing before the . on .png or .jpg
		ext := strings.Split(fh.Filename, ".")[1]
		// create sha for file name, returns a hash
		h := sha1.New()
		// Copy mf into h, copy takes a source and destination
		// func Copy(dst Writer, src Reader) (written int64, err error)
		io.Copy(h, mf)
		// h.Sum, is how you get the hash to work
		// %x prints as a hex value
		fname := fmt.Sprintf("%x", h.Sum(nil)) + "." + ext
		// create new file
		// os.Getwd, get current working dir
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}
		// Join current working dir + public + images, to cretate path
		path := filepath.Join(wd, "public", "images", fname)
		// Create a file using the path
		nf, err := os.Create(path)
		if err != nil {
			fmt.Println(err)
		}
		defer nf.Close()
		// copy
		// Reset the read/write head to the beginning of the file
		mf.Seek(0, 0)
		// Now copy contents of the multipart file to the new file nf
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
