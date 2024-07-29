package database

import (
	// "database/sql"
	"Forum/backend/structs"
	"database/sql"
	"strings"
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
ORDER BY
	posts.post_id DESC;`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []structs.Post
	for rows.Next() {
		var post structs.Post
		var categoryName string
		err := rows.Scan(&post.PostID, &post.UserID, &post.Dislike, &post.Like, &post.PostHeading, &post.Postdescription, &post.Username, &categoryName)
		if err != nil {
			return nil, err
		}
		post.CategoryName = strings.Split(categoryName, ",")
		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func GetPostById(id int) (structs.Post, error) {
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
WHERE
    posts.post_id = ?
GROUP BY 
    posts.post_id;`

	row := db.QueryRow(query, id)

	var post structs.Post
	var categoryName string
	err := row.Scan(&post.PostID, &post.UserID, &post.Dislike, &post.Like, &post.PostHeading, &post.Postdescription, &post.Username, &categoryName)
	if err != nil {
		if err == sql.ErrNoRows {
			return structs.Post{}, nil // No post found with the given ID
		}
		return structs.Post{}, err
	}
	post.CategoryName = strings.Split(categoryName, ", ")
	post.Comments, err = GetAllComments(id)
	if err != nil {
		return structs.Post{}, err
	}
	return post, nil
}
