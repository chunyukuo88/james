package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Woobler's in the yard!</h1>")
	fmt.Fprint(os.Stdout, "oh hello there, browser dev tools.\n")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<p>hi! To get in touch, email me at alexgochenour at gmail dot com</hp>")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>This is the FAQ page. But no one has asked any questions frequently so there's nothing else here.</h2>")
}

func galleryHandler(w http.ResponseWriter, r *http.Request) {
	galleryId := chi.URLParam(r, "galleryId")
	fmt.Fprintf(w, "<h1>Welcome to the %v gallery!</h1>", galleryId)
}

func pageNotFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Page not found!", http.StatusNotFound)
}

func main() {
	router := chi.NewRouter()
	router.Get("/", rootHandler)
	router.Post("/contact", contactHandler)
	router.Get("/faq", faqHandler)
	router.Get("/galleries/{galleryId}", galleryHandler)
	router.NotFound(pageNotFoundHandler)
	fmt.Println("Starting my rad server on port 3000...")
	http.ListenAndServe(":3000", router)
}
