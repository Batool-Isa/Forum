package main

import (
	database "Forum/backend"
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/error", errorHandler)
	http.HandleFunc("/create_post", createHandler)

	// Serve static files
	fs := http.FileServer(http.Dir("templates"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	db, errDb := sql.Open("sqlite3", "./forum.db")
	if errDb != nil {
		log.Fatal(errDb)
	}
	defer db.Close()

	// Create tables
	database.CreateTables(db)
	log.Println("Database setup complete")

	database.AddDummyData(db)
	database.ShowData(db)

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

func createHandler(w http.ResponseWriter, r *http.Request) { // New create handler
	renderTemplate(w, "create_post.html")
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
