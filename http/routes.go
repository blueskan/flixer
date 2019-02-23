package http

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

type FlixerRoutes interface {
	DefineRouteForRenderTemplate(route string, templatePath string) error
	DefineRouteForObtainInputs(route string)
}

type flixerRoutes struct {
	obtainInputsCh chan<- string
}

func (fr *flixerRoutes) DefineRouteForRenderTemplate(route string, templatePath string) error {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return fmt.Errorf("Error happened while opening the template file: %s", templatePath)
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	}

	http.HandleFunc(route, handler)
	return nil
}

func (fr *flixerRoutes) DefineRouteForObtainInputs(route string) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		postGuard(w, r)

		body, _ := ioutil.ReadAll(r.Body)
		fr.obtainInputsCh <- string(body)

		w.WriteHeader(http.StatusAccepted)
	}

	http.HandleFunc(route, handler)
}

func postGuard(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)

		fmt.Fprintf(w, "Only `POST` method accessible for obtaining inputs!")
	}
}

func NewFlixerRoutes(obtainInputsCh chan<- string) FlixerRoutes {
	return &flixerRoutes{
		obtainInputsCh: obtainInputsCh,
	}
}
