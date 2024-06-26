package main

import (
	"html/template"
	"log"
	"fmt"
	"net/http"
)
func main() {
    http.HandleFunc("/", indexHandler)
    http.HandleFunc("/register", registerHandler)
    http.HandleFunc("/login", loginHandler)
    http.HandleFunc("/error", errorHandler)

    // Serve static files
    fs := http.FileServer(http.Dir("templates"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    fmt.Println("Server started at http://localhost:8080/")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal("Error executing server at 8080", err)
    }
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html")
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "register.html")
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "login.html")
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "error.html")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	tmpl = "templates/" + tmpl
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		http.Error(w, "Template not found", http.StatusNotFound)
		return
	}
	t.Execute(w, nil)
}
