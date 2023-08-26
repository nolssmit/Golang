package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter,
		r *http.Request) {
		nb, err := fmt.Fprintf(w, "Hello Browser")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Bytes Written : %d", nb)
	})
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal((err))
}
/* 
To test it: Open a CMD-line window in folder containing main.go
Execute the command $ go run .
In the browser, in the address field:   localhost:8080
It will display:   Hello Browser
*/ 
