package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content", "text/html")
	fmt.Fprint(w, "<h1>Hello to my site!</h1>")
}

func contactFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content", "text/html")
	fmt.Fprint(w, "<h1>Contact me at pillai_amal@hotmail.com</h1>")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.HandleFunc("/contact", contactFunc)
	http.ListenAndServe(":3000", nil)
}
