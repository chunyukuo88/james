package main

import (
	"fmt"
	"net/http"
	"os"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Woobler's in the yard!</h1>")
	fmt.Fprint(os.Stdout, "oh hello there, browser dev tools.\n")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<p>hi! To get in touch, email me at alexgochenour at gmail dot com</hp>")
}

func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		rootHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		http.Error(w, "Where it is? I not knowing. But that page, it is not here.", http.StatusNotFound)
	}
}

func main() {
	http.HandleFunc("/", pathHandler)
	fmt.Println("Starting my rad server on port 3000...")
	http.ListenAndServe(":3000", nil)
}
