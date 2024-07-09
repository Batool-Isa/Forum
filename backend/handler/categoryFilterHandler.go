package handler

// import (
// 	"Forum/backend/database"
// 	"Forum/backend/structs"
// 	"fmt"
// 	"html/template"
// 	"log"
// 	"net/http"
// 	"strconv"
// )

// func CategoryFilterHandler(w http.ResponseWriter, r *http.Request) {
// 	categoryID, err := strconv.Atoi(r.URL.Query().Get("category_id"))
// 	if err != nil {
// 		log.Println("Error parsing category ID:", err)
// 		http.Error(w, "Bad Request", http.StatusBadRequest)
// 		return
// 	}

// 	posts, err := database.GetPostsByCategory(categoryID)
// 	if err != nil {
// 		log.Println("Error fetching posts:", err)
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	} 
// 	fmt.Println(posts)

// 	tmpl, err := template.ParseFiles("templates/index.html")
// 	if err != nil {
// 		log.Println("Error parsing template:", err)
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}

// 	data := struct {
// 		Posts []structs.Post
// 	}{
// 		Posts: posts,
// 	}

// 	err = tmpl.Execute(w, data)
// 	if err != nil {
// 		log.Println("Error executing template:", err)
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 	}


// }
