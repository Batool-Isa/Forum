package database

import (
	// "database/sql"
	"Forum/backend/structs"
)

// var db *sql.DB

func GetAllPosts() ([]structs.Post, error) {
	query := `
	SELECT 
    posts.post_id, 
    posts.user_id, 
    posts.dislike, 
    posts.like, 
    posts.post_heading, 
    posts.post_data, 
    users.username, 
    COALESCE(GROUP_CONCAT(categories.category_name), ', ', '') AS category_name	
FROM 
    posts
INNER JOIN 
    users ON posts.user_id = users.uid
LEFT JOIN 
    post_categories ON posts.post_id = post_categories.post_id
LEFT JOIN 
    categories ON post_categories.category_id = categories.category_id
GROUP BY 
    posts.post_id
LIMIT 100;`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	postMap := make(map[int]*structs.Post)
	for rows.Next() {
		var postID int
		var post structs.Post
		var categoryName string
		err := rows.Scan(&postID, &post.UserID, &post.Dislike, &post.Like, &post.PostHeading, &post.Postdescription, &post.Username, &categoryName)
		if err != nil {
			return nil, err
		}

		if existingPost, exists := postMap[postID]; exists {
			existingPost.CategoryName = categoryName
		} else {
			post.PostID = postID
			post.CategoryName = categoryName
			postMap[postID] = &post
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	posts := make([]structs.Post, 0, len(postMap))
	for _, post := range postMap {
		posts = append(posts, *post)
	}

	return posts, nil
}
