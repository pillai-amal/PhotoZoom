package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/go-chi/chi/v5"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content", "text/html")
	tmplPath := filepath.Join("templates", "home.gohtml")
	t, err := template.ParseFiles(tmplPath)
	if err != nil {
		log.Printf("%v error was occured on parsing", err)
		http.Error(w, "the error caused the program to quit", http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		log.Printf("%v error was occured on template execution", err)
		http.Error(w, "the error caused the program to quit", http.StatusInternalServerError)
		return
	}
}

func contactFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content", "text/html")
	tmplPath := filepath.Join("templates", "contact.gohtml")
	t, err := template.ParseFiles(tmplPath)
	if err != nil {
		log.Printf("%v error was occured on parsing", err)
		http.Error(w, "the error caused the program to quit", http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		log.Printf("%v error was occured on template execution", err)
		http.Error(w, "the error caused the program to quit", http.StatusInternalServerError)
		return
	}
}

func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		handlerFunc(w, r)
	case "/contact":
		contactFunc(w, r)
	default:
		http.Error(w, "Looks like you are on to page that does not exsists", http.StatusNotFound)
	}
}

func main() {
	r := chi.NewRouter()
	r.Get("/", pathHandler)
	fmt.Println("Go to http://localhost:3000/ in your browser")
	fmt.Println("Server started and is listening to the port 3000")
	http.ListenAndServe(":3000", r)
}
