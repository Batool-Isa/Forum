package handler

import (
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "login.html")
}