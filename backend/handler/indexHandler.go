package handler

import (
	"Forum/backend/database"
	"Forum/backend/structs"
	"html/template"
	"log"
	"net/http"
	// "structs"
)


// Define a map to store the category name to category ID mapping
var categoryMap = map[string]int{
	"all":         0, // 0 represents all categories
	"sports":      1,
	"technology":  2,
	"education":   3,
}


func IndexHandler(w http.ResponseWriter, r *http.Request) {
	posts, err := database.GetAllPosts()	
	if err != nil {
		log.Println("Error fetching posts:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	categorySelection := r.URL.Query().Get("filter-category")
	if categorySelection != "" && categorySelection != "all" {
		categoryID, ok := categoryMap[categorySelection]
		if !ok {
			// Handle invalid category selection
			return
		}
		posts, err = database.GetPostsByCategory(categoryID)
		if err != nil {
			log.Println("Error fetching posts by category:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
	// categorySelection := r.URL.Query().Get("filter-category")

	// if categorySelection != "" && categorySelection != "all"{
	// 	posts, err = database.GetPostsByCategory(categorySelection)
	// }

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
