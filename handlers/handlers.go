package handlers

import (
	"fmt"
	"net/http"

	"templofknor.com/services"
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
	st := &services.MyTier.STier
	vu := &services.MyTier.VeryUseful
	av := &services.MyTier.Average
	bp := &services.MyTier.Bupkis
	dw := &services.MyTier.DogWater
	if r.Method == "GET" {
		fmt.Println("Add image to tier.")
		fmt.Println(r.Header.Get("Hx-Target"))
		switch lang := r.Header.Get("Hx-Target"); lang {
		case "c.png":
			fmt.Println("S Tier")
			*st = append(*st, services.TierItem{Location: "c.png"})
			fmt.Println(st)
		case "cpp.png":
			fmt.Println("Very Useful")
			*vu = append(*vu, services.TierItem{Location: "cpp.png"})
			fmt.Println(vu)
		case "csharp.png":
			fmt.Println("Very Useful")
			*vu = append(*vu, services.TierItem{Location: "csharp.png"})
			fmt.Println(vu)
		case "golang.png":
			fmt.Println("S Tier")
			*st = append(*st, services.TierItem{Location: "golang.png"})
			fmt.Println(st)
		case "haskell.png":
			fmt.Println("Average")
		case "java.png":
			fmt.Println("Dog Water")
			*dw = append(*dw, services.TierItem{Location: "java.png"})
			fmt.Println(dw)
		case "rust.png":
			fmt.Println("Bupkis")
			*bp = append(*bp, services.TierItem{Location: "rust.png"})
			fmt.Println(bp)
		case "zig.png":
			fmt.Println("Average")
			*av = append(*dw, services.TierItem{Location: "rust.png"})
			fmt.Println(av)
		default:
			fmt.Println("Average")

		}

	} else if r.Method == "POST" {
		fmt.Println("Just do a GET.")
	}
}
