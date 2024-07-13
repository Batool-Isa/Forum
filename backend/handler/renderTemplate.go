package handler

import (
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmpl string, errors ...map[string]string) {
	t, err := template.ParseFiles("templates/" + tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(errors) > 0 {
		errs := struct {
			Errors map[string]string
		}{
			Errors: errors[0],
		}
		err = t.Execute(w, errs)
	} else {
		err = t.Execute(w, nil)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
