package handler

import (
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	templateFile := "error.html"
	data := struct {
		StatusCode  int
		StatusText  string
		Description string
	}{
		StatusCode:  status,
		StatusText:  http.StatusText(status),
		Description: getErrorDescription(status),
	}
	templateData := TemplateData{
		Data: data,
		Errors: nil,
	}

	RenderTemplate(w, r, templateFile, templateData)
}

func getErrorDescription(status int) string {
	switch status {
	case http.StatusBadRequest:
		return "Sorry, the request is invalid or incomplete."
	case http.StatusNotFound:
		return "Sorry, the page you are looking for might be missing or the URL is incorrect."
	case http.StatusInternalServerError:
		return "Sorry, something went wrong on our end. We are working to fix the issue."
	default:
		return "An unexpected error occurred."
	}
}
