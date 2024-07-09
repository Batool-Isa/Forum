package database

import (
	"Forum/backend/structs"
	"strings"
)

func GetAllPosts() ([]structs.Post, error) {
	rows, err := db.Query("SELECT posts.post_heading, posts.post_data, (SELECT count() FROM likes WHERE likes.post_id = posts.post_id) AS like_count,(SELECT count() FROM dislikes WHERE dislikes.post_id = posts.post_id) AS dislike_count, GROUP_CONCAT(categories.category_name) AS categories, user_id, username, posts.post_id FROM posts LEFT JOIN post_categories ON post_categories.post_id = posts.post_id LEFT JOIN categories ON categories.category_id = post_categories.category_id INNER JOIN users ON users.uid= posts.user_id GROUP BY posts.post_id;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
 
	var posts []structs.Post
	for rows.Next() {
		var post structs.Post
		var categories string
		err := rows.Scan(&post.PostHeading, &post.Postdescription, &post.Dislike, &post.Like, &categories, &post.UserID, &post.Username, &post.PostID)
		if err != nil {
			return nil, err
		}
		post.Categories = strings.Split(categories, ",")
		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}
