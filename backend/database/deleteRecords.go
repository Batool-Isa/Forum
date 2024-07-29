package database

import (
	// "database/sql"

	"log"

	_ "github.com/mattn/go-sqlite3"
)

func DeleteLike(postId int, userId int) {
	stmt1, err := db.Prepare("DELETE FROM likes WHERE post_id = ? AND user_id= ?")
	if err != nil {
		log.Fatalln(err)
	}
	_, err = stmt1.Exec(postId, userId)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("Deleted like")
	}
}

func DeleteDislike(postId int, userId int) {
	stmt1, err := db.Prepare("DELETE FROM dislikes WHERE post_id = ? AND user_id= ?")
	if err != nil {
		log.Fatalln(err)
	}
	_, err = stmt1.Exec(postId, userId)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("Deleted dislike")
	}
}

func DeleteCommentLike(CommentID int, userId int) {
	stmt1, err := db.Prepare("DELETE FROM likeComment WHERE comment_id = ? AND user_id = ?")
	if err != nil {
		log.Fatalln(err)
	}
	_, err = stmt1.Exec(CommentID, userId)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("Deleted comment like")
	}
}

func DeleteCommentDislike(CommentID int, userId int) {
	stmt1, err := db.Prepare("DELETE FROM dislikeComment WHERE comment_id = ? AND user_id = ?")
	if err != nil {
		log.Fatalln(err)
	}
	_, err = stmt1.Exec(CommentID, userId)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("Deleted comment dislike")
	}
}