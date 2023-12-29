package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
)

func HandlePrint(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Println("This is a GET request")
		w.Write([]byte("This is a GET request."))
	} else if r.Method == "POST" {
		fmt.Println("This is a POST request.")
		w.Write(Sum(1, 2))
	}
}

func main() {
	component := hello("John")
	componentTwo := root("Gerald")

	http.Handle("/", templ.Handler(componentTwo))
	http.Handle("/hello", templ.Handler(component))
	http.HandleFunc("/print", HandlePrint)

	fmt.Println("Listening on port 3333")
	http.ListenAndServe(":3333", nil)
}
