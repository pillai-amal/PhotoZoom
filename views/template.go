package views

import (
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
)

func ParseFS(fs fs.FS, pattern string) (Template, error) {
	t, err := template.ParseFS(fs, pattern)
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
