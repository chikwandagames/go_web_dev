package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	// tpl, is a pointer to a template
	// tpl, is like a container for all our templates
	// Step 1. parse file
	tpl, err := template.ParseFiles("tpl.html")

	if err != nil {
		log.Fatalln(err)
	}

	// Create new file
	nf, err := os.Create("index.html")
	if err != nil {
		log.Println("error creating files", err)
	}

	// Close the file after program exits
	defer nf.Close()

	// Step 2, execute file
	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatalln(err)
	}

}
