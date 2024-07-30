package utils

import (
	//"Forum/backend/handler"
	"html/template"
	"net/http"
)

type TemplateData struct {
	Errors map[string]string
	Data   interface{}
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, data interface{}, errors ...map[string]string) {
	t, err := template.ParseFiles("templates/" + tmpl)
	if err != nil {
		ErrorHandler(w,r,http.StatusNotFound)
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	td := TemplateData{
		Data: data,
	}

	// Ensure errors are properly assigned
	if len(errors) > 0 {
		td.Errors = errors[0]
	} else {
		td.Errors = nil
	}
	
	err = t.Execute(w, td)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		//http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
