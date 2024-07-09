package database

import (
	"database/sql"
	"fmt"
	"Forum/backend/structs"
)

func GetPostsByCategory(categoryID int) ([]structs.Post, error) {
	query := `
        SELECT p.post_id, p.post_heading, p.post_data, p.user_id, p.dislike, p.like
        FROM posts p
        JOIN post_categories pc ON p.post_id = pc.post_id
        WHERE pc.category_id = ?
    `

	rows, err := db.Query(query, categoryID)
	if err != nil {
		return nil, fmt.Errorf("query error: %v", err)
	}
	defer rows.Close()

	var posts []structs.Post
	for rows.Next() {
		var post structs.Post
		var userID sql.NullInt64
		var dislike sql.NullInt64
		var like sql.NullInt64

		// Scan NULL-able fields
		err := rows.Scan(&post.PostID, &post.PostHeading, &post.Postdescription, &userID, &dislike, &like)
		if err != nil {
			return nil, fmt.Errorf("scan error: %v", err)
		}

		// Check for NULL values and convert to appropriate types
		if userID.Valid {
			post.UserID = int(userID.Int64)
		}
		if dislike.Valid {
			post.Dislike = int(dislike.Int64)
		}
		if like.Valid {
			post.Like = int(like.Int64)
		}

		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row error: %v", err)
	}

	return posts, nil
}
