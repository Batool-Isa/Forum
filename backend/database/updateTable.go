package database

import (
	// "database/sql"

	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func UpdatePost(postId int) {
	fmt.Printf("Post ID is %v\n", postId)
	stmt1, err := db.Prepare("UPDATE posts SET 'like'= (SELECT count(*) FROM likes where post_id=?) where post_id=?;")
	if err != nil {
		log.Fatalln(err)
	}
	_, err = stmt1.Exec(postId, postId)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("Updated likes")
	}

	stmt2, err := db.Prepare("UPDATE posts SET dislike=(SELECT count(*) FROM dislikes where post_id=?) where post_id=?;")
	if err != nil {
		log.Fatalln(err)
	}
	_, err = stmt2.Exec(postId, postId)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("Updated dislikes")
	}
}
