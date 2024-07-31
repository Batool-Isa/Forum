package handler

import (
	"net/http"

	"Forum/backend/database"
	"Forum/backend/structs"
	"Forum/backend/utils"

	"Forum/backend/middleware"
)

// Define a map to store the category name to category ID mapping
var categoryMap = map[string]int{
	"all":           0, // 0 represents all categories
	"sports":        1,
	"technology":    2,
	"education":     3,
	"health":        5,
	"entertainment": 6,
	"travel":        7,
	"finance":       8,
	"culture":       9,
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		utils.ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	posts, err := database.GetAllPosts()
	if err != nil {
		utils.ErrorHandler(w, r, http.StatusInternalServerError)
		//http.Error(w, "Internal Server Error", http.StatusInternalServerError)
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
			utils.ErrorHandler(w, r, http.StatusInternalServerError)

			//http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}

	session := middleware.FromContext(r.Context())

	myPosts := r.URL.Query().Get("filter-mypost")
	likedepost := r.URL.Query().Get("filter-likedpost")

	if session != nil {
		userId := session.UserID
		if myPosts == "true" {
			posts, err = database.GetPostByUserID(userId)
			if err != nil {
				utils.ErrorHandler(w, r, http.StatusInternalServerError)

				//http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
		} else if likedepost == "true" {
			posts, err = database.GetLikedPost(userId)
			if err != nil {
				utils.ErrorHandler(w, r, http.StatusInternalServerError)

				//http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

		}
	}

	//  posts = database.GetPostByUserID(userId)

	data := struct {
		Posts   []structs.Post
		Session *structs.Session
	}{
		Posts:   posts,
		Session: session,
	}

	// err = tmpl.Execute(w, data)
	// if err != nil {
	// 	log.Println("Error executing template:", err)
	// 	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	// }
	utils.RenderTemplate(w, r, "index.html", data)

}
