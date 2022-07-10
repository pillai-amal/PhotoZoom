package views

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func Parse(filepath string) (Template, error) {
	t, err := template.ParseFiles(filepath)
	if err != nil {
		return Template{}, fmt.Errorf("%w Error occoured during parsing", err)
	}
	return Template{
		htmlTmpl: t,
	}, nil
}

type Template struct {
	htmlTmpl *template.Template
}

func (t Template) Execute(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content", "text/html")
	err := t.htmlTmpl.Execute(w, data)
	if err != nil {
		log.Printf("%v error was occured on template execution", err)
		http.Error(w, "the error caused the program to quit", http.StatusInternalServerError)
		return
	}
}
