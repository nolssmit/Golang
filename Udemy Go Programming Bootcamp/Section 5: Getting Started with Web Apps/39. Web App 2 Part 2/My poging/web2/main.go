package main

import (
	"errors"
	"fmt"
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

func addHandler(writer http.ResponseWriter,
	request *http.Request) {
		write(writer, "Hello internet add-sum calculator\n")
		sum := getSum(5, 4)
		output := fmt.Sprintf("5 + 4 = %d\n", sum)
		write(writer, output)
	}

func getSum(x, y int) int {
	return x + y
}

func divideHandler(writer http.ResponseWriter,
	request *http.Request) {
		v, err := getQuotient(5, 4)
		if err != nil {
			write(writer, "cant't divide by zero\n")
		} else {
			output := fmt.Sprintf("5 / 4 = %.2f\n", v)
			write(writer, output)
		}
	}

func getQuotient(x, y float32) (float32, error) {
	if y == 0 {
		err := errors.New("cannot divide bt zero")
		return 0, err
	} else {
		return (x/y), nil
	}
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/getsum", addHandler)	
	http.HandleFunc("/divide", divideHandler)	

	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
// Test with localhost:8080/get sum or localhost:8080/divide