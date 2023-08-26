package main

import (
	"log"
	"net/http"
	"text/template"
)

func write(writer http.ResponseWriter,
	msg string) {
	_, err := writer.Write([]byte(msg)) // convert string message to array of bytes
	if err != nil {
		log.Fatal(err)
	}
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, err := template.ParseFiles("./templates/" + tmpl)
	errorCheck(err)
	err = parsedTemplate.Execute(w, nil)
	errorCheck(err)
}

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func homeHandler(w http.ResponseWriter,
	request *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

func main() {
	http.HandleFunc("/", homeHandler)

	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}

// If you get an error message: listen tcp 127.0.0.1:8080: bind: address already in use
// In a terminal window: $ lsof -i :8080    and then: $ kill <PID>
 

