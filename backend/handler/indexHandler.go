package handler

import (
	"Forum/backend/database"
	"Forum/backend/structs"
	"html/template"
	"log"
	"net/http"
	// "structs"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	posts, err := database.GetAllPosts()
	if err != nil {
		log.Println("Error fetching posts:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Println("Error parsing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	data := struct {
		Posts []structs.Post
	}{
		Posts: posts,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println("Error executing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	// RenderTemplate(w, "index.html")
}
