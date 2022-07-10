package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content", "text/html")
	fmt.Fprint(w, "<h1>Hello to my site!</h1>")
}

func contactFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content", "text/html")
	fmt.Fprint(w, "<h1>Contact me at pillai_amal@hotmail.com</h1>")
}

func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		handlerFunc(w, r)
		break
	case "/contact":
		contactFunc(w, r)
		break
	default:
		http.Error(w, "Looks like you are on to page that does not exsists", http.StatusNotFound)
	}
}

func main() {
	r := chi.NewRouter()
	fmt.Println("Go to http://localhost:3000/ in your browser")
	fmt.Println("Server started and is listening to the port 3000")
	http.ListenAndServe(":3000", r)
}
