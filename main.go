package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Woobler's in the yard!</h1>")
	fmt.Fprint(os.Stdout, "oh hello there, browser dev tools.\n")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting my rad server on port 3000...")
	http.ListenAndServe(":3000", nil)
}
