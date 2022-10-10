package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("templates/home.gohtml")
	if err != nil {
		handleTemplateError(err, "parse", w)
	}
	err = template.Execute(w, nil)
	if err != nil {
		handleTemplateError(err, "execution", w)
	}
}

func handleTemplateError(err error, errorType string, w http.ResponseWriter) {
	if errorType == "parse" {
		log.Printf("Error parsing the template: %v", err)
		http.Error(w, "There was an error parsing the template.", http.StatusInternalServerError)
		return
	} else {
		log.Printf("Error executing the template: %v", err)
		http.Error(w, "Error executing the template.", http.StatusInternalServerError)
		return
	}
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
	router.Use(middleware.Logger)
	router.Get("/", homeHandler)
	router.Post("/contact", contactHandler)
	router.Get("/faq", faqHandler)
	router.Get("/galleries/{galleryId}", galleryHandler)
	router.NotFound(pageNotFoundHandler)
	fmt.Println("Starting my rad server on port 3000...")
	http.ListenAndServe(":3000", router)
}
