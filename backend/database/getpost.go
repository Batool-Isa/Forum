package database

import "Forum/backend/structs"

func GetAllPosts() ([]structs.Post, error) {
	rows, err := db.Query("SELECT post_id, user_id, dislike, like, post_heading, post_data FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []structs.Post
	for rows.Next() {
		var post structs.Post
		err := rows.Scan(&post.PostID, &post.UserID, &post.Dislike, &post.Like, &post.PostHeading, &post.Postdescription)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}
