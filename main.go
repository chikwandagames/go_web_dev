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
	tpl, err := template.ParseFiles("three.wizzy")

	if err != nil {
		log.Fatalln(err)
	}

	// Step 2, execute file
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	// Add more files to the tamplate container
	tpl, err = tpl.ParseFiles("one.wizzy", "two.wizzy")
	if err != nil {
		log.Fatalln(err)
	}

	// You can specify which template to execute
	err = tpl.ExecuteTemplate(os.Stdout, "one.wizzy", nil)
	if err != nil {
		log.Fatalln(err)
	}

}
