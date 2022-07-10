package controllers

import (
	"net/http"

	"phtozoom.com/m/views"
)

func StaticHandler(t views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t.Execute(w, nil)
	}
}

func FAQHandler(t views.Template) http.HandlerFunc {
	questions := []struct {
		Question string
		Answer   string
	}{
		{
			Question: "What is this website About",
			Answer:   "This is inted to learn web dev in Golang",
		},
		{
			Question: "Who is the author of this project",
			Answer:   "@pillai_amal is the author",
		},
		{
			Question: "How can I contact the author",
			Answer:   "Amal can be contacted on github at @pillai_amal",
		},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		t.Execute(w, questions)
	}
}
