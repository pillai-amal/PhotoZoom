package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"phtozoom.com/m/controllers"
	"phtozoom.com/m/templates"
	"phtozoom.com/m/views"
)

func main() {
	r := chi.NewRouter()
	tmpl, err := views.ParseFS(templates.FS, "home.gohtml", "tailwind.gohtml")
	if err != nil {
		panic(err)
	}
	r.Get("/", controllers.StaticHandler(tmpl))
	//
	tmpl, err = views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml")
	if err != nil {
		panic(err)
	}
	r.Get("/contact", controllers.StaticHandler(tmpl))

	tmpl, err = views.ParseFS(templates.FS, "faq.gohtml", "tailwind.gohtml")
	if err != nil {
		panic(err)
	}
	r.Get("/faq", controllers.FAQHandler(tmpl))
	fmt.Println("Go to http://localhost:3000/ in your browser")
	fmt.Println("Server started and is listening to the port 3000")
	http.ListenAndServe(":3000", r)
}
