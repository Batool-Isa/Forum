package database

import (
	"Forum/backend/structs"
)

func GetAllComments(postID int) ([]structs.Comment, error) {
	query := `
    SELECT
        comments.comment_id,
        comments.user_id,
        comments.post_id,
        comments.comment,
        users.username
    FROM
        comments
    INNER JOIN
        users ON comments.user_id = users.uid
    WHERE
        comments.post_id = ?`

	rows, err := db.Query(query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comms []structs.Comment
	for rows.Next() {
		var comm structs.Comment
		err := rows.Scan(&comm.CommentID, &comm.UserID, &comm.PostID, &comm.Text, &comm.UserName)
		if err != nil {
			return nil, err
		}
		comms = append(comms, comm)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return comms, nil

}
