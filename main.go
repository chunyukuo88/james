package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	filePath := filepath.Join("templates", "home.gohtml")
	pageHandler(w, r, filePath)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	filePath := filepath.Join("templates", "contact.gohtml")
	pageHandler(w, r, filePath)
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	filePath := filepath.Join("templates", "faq.gohtml")
	pageHandler(w, r, filePath)
}

func galleryHandler(w http.ResponseWriter, r *http.Request) {
	// galleryId := chi.URLParam(r, "galleryId") TODO: Add this back in once I figure out how dynamic URLs work with templating. Slugs?
	filePath := filepath.Join("templates", "gallery.gohtml")
	pageHandler(w, r, filePath)
}

func pageHandler(w http.ResponseWriter, r *http.Request, filePath string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	template, err := template.ParseFiles(filePath)
	handlePossibleTemplateError(err, "parse", w)
	err = template.Execute(w, nil)
	handlePossibleTemplateError(err, "execution", w)
}

func handlePossibleTemplateError(err error, errorType string, w http.ResponseWriter) {
	if err != nil {
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
}

func pageNotFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Page not found!", http.StatusNotFound)
}

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", homeHandler)
	router.Get("/contact", contactHandler)
	router.Get("/faq", faqHandler)
	router.Get("/galleries/{galleryId}", galleryHandler)
	router.NotFound(pageNotFoundHandler)
	fmt.Println("Starting my rad server on port 3000...")
	http.ListenAndServe(":3000", router)
}
