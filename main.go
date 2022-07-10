package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"phtozoom.com/m/controllers"
	"phtozoom.com/m/views"
)

func main() {
	r := chi.NewRouter()
	tmpl, err := views.Parse(filepath.Join("templates", "home.gohtml"))
	if err != nil {
		panic(err)
	}
	r.Get("/", controllers.StaticHandler(tmpl))
	//
	tmpl, err = views.Parse(filepath.Join("templates", "contact.gohtml"))
	if err != nil {
		panic(err)
	}
	r.Get("/contact", controllers.StaticHandler(tmpl))
	fmt.Println("Go to http://localhost:3000/ in your browser")
	fmt.Println("Server started and is listening to the port 3000")
	http.ListenAndServe(":3000", r)
}
