package handler

import (
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "error.html", nil)
}
