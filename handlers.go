package main

import (
	"fmt"
	"net/http"
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

func AddToTier(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Println(tiers, options)
	} else if r.Method == "POST" {
		fmt.Println("This is a POST request.")
	}
}
