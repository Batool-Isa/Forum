package database

import (
	"Forum/backend/structs"
	"fmt"
	"strings"
)

func GetPostByUserID(userID int) ([]structs.Post, error) {
	query := `SELECT 
		p.post_id,
		p.user_id,
		p.dislike,
		p.like,
		p.post_heading,
		p.post_data,
		u.username,
		COALESCE(GROUP_CONCAT(c.category_name), ', ', '') AS category_name
	FROM
		posts p
	INNER JOIN
		users u ON p.user_id = u.uid
	LEFT JOIN
		post_categories pc ON p.post_id = pc.post_id
	LEFT JOIN
		categories c ON pc.category_id = c.category_id
	WHERE
		p.user_id = ?
	GROUP BY
		p.post_id
	ORDER BY
		p.post_id DESC;

	`
	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("query error: %v", err)
	}
	defer rows.Close()

	var posts []structs.Post
	for rows.Next() {
		var post structs.Post
		var categoryName string
		err := rows.Scan(&post.PostID, &post.UserID, &post.Dislike, &post.Like, &post.PostHeading, &post.Postdescription, &post.Username, &categoryName)
		if err != nil {
			return nil, fmt.Errorf("scan error: %v", err)
		}
		post.CategoryName = strings.Split(categoryName, ",")
		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row error: %v", err)
	}

	return posts, nil
}