package handler

import (
	"fmt"
	"log"
	"net/http"

	"Forum/backend/database"
	"Forum/backend/structs"


	"Forum/backend/middleware"
)

// Define a map to store the category name to category ID mapping
var categoryMap = map[string]int{
	"all":        0, // 0 represents all categories
	"sports":     1,
	"technology": 2,
	"education":  3,
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	posts, err := database.GetAllPosts()
	if err != nil {
		log.Println("Error fetching posts:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	categorySelection := r.URL.Query().Get("filter-category")
	fmt.Println("Category Selection: ", categorySelection)
	if categorySelection != "" && categorySelection != "all" {
		categoryID, ok := categoryMap[categorySelection]
		fmt.Println("Category ID: ", categoryID)
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


	session := middleware.FromContext(r.Context())

	myPosts := r.URL.Query().Get("filter-mypost")
    likedepost := r.URL.Query().Get("filter-likedpost")
	fmt.Println("My Posts: ", myPosts)
	fmt.Println("Liked Posts: ", likedepost)

if session != nil{  
		userId := session.UserID  
		fmt.Println("User ID: ", userId)
   if myPosts == "true" {
	posts, err = database.GetPostByUserID(userId)
	if err != nil {
		log.Println("Error fetching posts:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}else if likedepost == "true" {
	posts, err = database.GetLikedPost(userId)
	fmt.Println("Liked Posts: ", posts)
	if err != nil {
		log.Println("Error fetching posts:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

     }
	}

    //  posts = database.GetPostByUserID(userId)

	data := struct {
		Posts   []structs.Post
		Session *structs.Session
	}{
		Posts: posts,
		Session: session,
	}

	// err = tmpl.Execute(w, data)
	// if err != nil {
	// 	log.Println("Error executing template:", err)
	// 	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	// }

	RenderTemplate(w, "index.html", data)
}
