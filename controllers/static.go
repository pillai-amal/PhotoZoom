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
