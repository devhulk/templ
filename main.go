package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"templofknor.com/components"
	"templofknor.com/handlers"
	"templofknor.com/services"
)

func main() {
	//componentTwo := root("Gerald")
	t := services.MyTier
	HomePage := components.HomePage(t, &t.Options)

	//http.Handle("/", templ.Handler(componentTwo))
	http.Handle("/", templ.Handler(HomePage))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/refresh", handlers.Refresh)
	http.HandleFunc("/print", handlers.HandlePrint)
	http.HandleFunc("/addToTier", handlers.AddToTier)

	fmt.Println("Listening on port 3333")
	http.ListenAndServe(":3333", nil)
}
