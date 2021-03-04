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

	// Print tpl.html data to the terminal
	// Step 2, execute file
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

}
