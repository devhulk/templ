package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
)

var options = &map[string]string{
	"C":          "c.png",
	"Golang":     "golang.png",
	"Python":     "python.png",
	"Java":       "java.png",
	"C++":        "cpp.png",
	"C#":         "csharp.png",
	"Haskell":    "haskell.png",
	"Zig":        "zig.png",
	"Rust":       "rust.png",
	"Javascript": "js.png",
}

var tiers = &map[string][]string{
	"S Tier":      []string{},
	"Very Useful": []string{},
	"Average":     []string{},
	"Bupkis":      []string{},
	"Dog Water":   []string{},
}

func main() {
	componentTwo := root("Gerald")
	tierList := TierList(tiers, options)

	http.Handle("/", templ.Handler(componentTwo))
	http.Handle("/tier", templ.Handler(tierList))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("img"))))
	http.HandleFunc("/print", HandlePrint)
	http.HandleFunc("/addToTier", AddToTier)

	fmt.Println("Listening on port 3333")
	http.ListenAndServe(":3333", nil)
}
