package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"phtozoom.com/m/views"
)

func excecuteTemplate(w http.ResponseWriter, tmplPath string) {
	t, err := views.Parse(tmplPath)
	if err != nil {
		log.Printf("%v error occured during the parsing", err)
		http.Error(w, "There was a error parsing the template", http.StatusInternalServerError)
	}
	t.Execute(w, nil)
}
func homeFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content", "text/html")
	tmplPath := filepath.Join("templates", "home.gohtml")
	excecuteTemplate(w, tmplPath)
}

func contactFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content", "text/html")
	tmplPath := filepath.Join("templates", "contact.gohtml")
	excecuteTemplate(w, tmplPath)
}

func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeFunc(w, r)
	case "/contact":
		contactFunc(w, r)
	default:
		http.Error(w, "Looks like you are on to page that does not exsists", http.StatusNotFound)
	}
}

func main() {
	r := chi.NewRouter()
	r.Get("/", pathHandler)
	r.Get("/contact", contactFunc)
	fmt.Println("Go to http://localhost:3000/ in your browser")
	fmt.Println("Server started and is listening to the port 3000")
	http.ListenAndServe(":3000", r)
}
