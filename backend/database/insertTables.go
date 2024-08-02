package database

import (
	// "database/sql"

	"log"
	"time"
	"Forum/backend/utils"
	_ "github.com/mattn/go-sqlite3"
)



func InsertUser(username string, password string, email string) error {
	InputErr := utils.ValidateInput(map[string]string{"username": username, "password": password, "email": email})
	if InputErr != nil {
		return InputErr
	}
	stmt, err := db.Prepare("INSERT INTO users(username,password,email) values(?,?,?)")
	if err != nil {
		// return err
		log.Println(err)
		return err
	}
	_, err = stmt.Exec(username, password, email)
	if err != nil {
		log.Println(err)
		// return err
		return err
	}
	return nil
}

func InsertCategories(category_name string) error {
	categoryArray, err := GetCategories()
	if err != nil {
		log.Println(err)
		return err
	}
	catErr := utils.ValidateCategory(category_name, categoryArray)
	if catErr != nil {
		return catErr
	}
	stmt, err := db.Prepare("INSERT INTO categories(category_name) values(?)")
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = stmt.Exec(category_name)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
func InsertPost(user_id int, post_heading string, post_data string, categoryName []string) error {
	postErr := utils.ValidatePost(post_heading, post_data, categoryName)
	if postErr != nil {
		return postErr
	}
	stmt, err := db.Prepare("INSERT INTO posts(user_id, post_heading, post_data) VALUES (?, ?, ?)")
	if err != nil {
		log.Println(err)
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(user_id, post_heading, post_data)
	if err != nil {
		log.Println(err)
		return err
	}

	postID, err := res.LastInsertId()
	if err != nil {
		log.Println(err)
		return err
	}

	for _, categoryName := range categoryName {

		var categoryID int
		err := db.QueryRow("SELECT category_id FROM categories WHERE category_name = ?", categoryName).Scan(&categoryID)
		if err != nil {
			log.Println(err)
			return err
		}

		InsertPostCategories(int(postID), categoryID)
	}
	return nil
}

func InsertComment(comment string, user_id int, postId int) error {
	CommentErr := utils.ValidateInput(map[string]string{"comment": comment})
	if CommentErr != nil {
		return CommentErr
	}
	stmt, err := db.Prepare("INSERT INTO comments(comment, user_id, post_id) values (?, ?, ?)")
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = stmt.Exec(comment, user_id, postId)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func InsertPostCategories(post_id int, category_id int) error {
	postCatErr := utils.ValidateInput(map[string]string{"post_id": string(post_id), "category_id": string(category_id)})
	if postCatErr != nil {
		return postCatErr
	}
	stmt, err := db.Prepare("INSERT INTO post_categories(post_id, category_id) values (?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(post_id, category_id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func InsertLikes(post_id int, user_id int) error {
	likeserr := utils.ValidateInput(map[string]string{"post_id": string(post_id), "user_id": string(user_id)})	
	if likeserr != nil {
		return likeserr
	}
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM likes WHERE post_id = ? AND user_id = ?", post_id, user_id).Scan(&count)
	if err != nil {
		log.Println(err)
		return err
	}

	if count == 1 {
		// If the like exists, delete it (unlike)
		stmt, err := db.Prepare("DELETE FROM likes WHERE post_id = ? AND user_id = ?")
		if err != nil {
			log.Println(err)
			return err
		}
		_, err = stmt.Exec(post_id, user_id)
		if err != nil {
			return err
		}
	} else {
		// If the like does not exist, insert it (like)
		stmt, err := db.Prepare("INSERT INTO likes(post_id, user_id) values (?, ?)")
		if err != nil {
			log.Println(err)
			return err
		}
		_, err = stmt.Exec(post_id, user_id)
		if err != nil {
			log.Println(err)
			return err
		}
	}

	return nil
}

func InsertDislikes(post_id int, user_id int) error {
	dislikeserr := utils.ValidateInput(map[string]string{"post_id": string(post_id), "user_id": string(user_id)})
	if dislikeserr != nil {
		return dislikeserr
	}
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM dislikes WHERE post_id = ? AND user_id = ?", post_id, user_id).Scan(&count)
	if err != nil {
		log.Println(err)
		return err
	}

	if count == 1 {
		// If the like exists, delete it (unlike)
		stmt, err := db.Prepare("DELETE FROM dislikes WHERE post_id = ? AND user_id = ?")
		if err != nil {
			log.Println(err)
			return err
		}
		_, err = stmt.Exec(post_id, user_id)
		if err != nil {
			log.Println(err)
			return err
		}
	} else {
		// If the like does not exist, insert it (like)
		stmt, err := db.Prepare("INSERT INTO dislikes(post_id, user_id) values (?, ?)")
		if err != nil {
			log.Println(err)
			return err
		}
		_, err = stmt.Exec(post_id, user_id)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

func InsertCommentLikes(comment_id int, user_id int) error {
	commentLikeErr := utils.ValidateInput(map[string]string{"comment_id": string(comment_id), "user_id": string(user_id)})
	if commentLikeErr != nil {
		return commentLikeErr
	}
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM likeComment WHERE comment_id = ? AND user_id = ?", comment_id, user_id).Scan(&count)
	if err != nil {
		log.Println(err)
		return err
	}

	if count == 1 {
		// If the like exists, delete it (unlike)
		stmt, err := db.Prepare("DELETE FROM likeComment WHERE comment_id = ? AND user_id = ?")
		if err != nil {
			log.Println(err)
			return err
		}
		_, err = stmt.Exec(comment_id, user_id)
		if err != nil {
			log.Println(err)
			return err
		}
	} else {
		// If the like does not exist, insert it (like)
		stmt, err := db.Prepare("INSERT INTO likeComment(comment_id, user_id) values (?, ?)")
		if err != nil {
			log.Println(err)
			return err
		}
		_, err = stmt.Exec(comment_id, user_id)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

func InsertCommentDislikes(comment_id int, user_id int) error {
	commentDislikeErr := utils.ValidateInput(map[string]string{"comment_id": string(comment_id), "user_id": string(user_id)})
	if commentDislikeErr != nil {
		return commentDislikeErr
	}
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM dislikeComment WHERE comment_id = ? AND user_id = ?", comment_id, user_id).Scan(&count)
	if err != nil {
		log.Println(err)
		return err
	}

	if count == 1 {
		// If the like exists, delete it (unlike)
		stmt, err := db.Prepare("DELETE FROM dislikeComment WHERE comment_id = ? AND user_id = ?")
		if err != nil {
			log.Println(err)
			return err
		}
		_, err = stmt.Exec(comment_id, user_id)
		if err != nil {
			log.Println(err)
			return err
		}
	} else {
		// If the like does not exist, insert it (like)
		stmt, err := db.Prepare("INSERT INTO dislikeComment(comment_id, user_id) values (?, ?)")
		if err != nil {
			log.Println(err)
			return err
		}
		_, err = stmt.Exec(comment_id, user_id)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

func InsertSession(session string, user_id int) error {
	sessionErr := utils.ValidateInput(map[string]string{"session": session, "user_id": string(user_id)})
	if sessionErr != nil {
		return sessionErr
	}
	stmt, err := db.Prepare("INSERT INTO sessions(session, user_id, timestamp) values (?, ?, ?)")
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = stmt.Exec(session, user_id, time.Now().Add(12*time.Hour))
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func AddDummyData() {

	// InsertUser(db, "zhashim", "12345678", "zahra@gmail.com")
	// InsertUser(db, "zee", "12345678", "z@gmail.com")
	// InsertUser(db, "zozo", "12345678", "zozo@gmail.com")

	// InsertCategories("Sports")
	// InsertCategories("Technology")
	// InsertCategories("Education")
	// InsertCategories("Health")
	// InsertCategories("Entertainment")
	// InsertCategories("Travel")
	// InsertCategories("Finance")
	// InsertCategories("Culture")

	// InsertPost(db, 1, "Advancements in AI Technology", "Explore the latest advancements in artificial intelligence and its applications.")
	// InsertPost(db, 1, "Online Learning Platforms", "Discover the best online learning platforms to enhance your skills and knowledge.")
	// InsertPost(db, 1, "Testing", "This is a test to test the test")

	// InsertPostCategories(db,2,3)
	// InsertPostCategories(db,1,2)
	// InsertPostCategories(db,3,3)
	// InsertPostCategories(db,3,4)
	// InsertPostCategories(9,4)
	// InsertPostCategories(10,3)
	// InsertPostCategories(10,1)

	// InsertComment(db, "Helpful info",2,2)
	// InsertComment(db, "Weirddd Stuff",3,2)
	// InsertComment(db, "Nice",3,3)

}

// func ShowData() {
// 	rows, err := db.Query("SELECT * FROM posts", "search")
// 	if err != nil {
// 		return err
// 	}
// 	defer rows.Close()
// 	fmt.Println("DATA FOR POST")
// 	for rows.Next() {
// 		var post_id int
// 		var user_id int
// 		var dislike int
// 		var like int
// 		var post_heading string
// 		var post_data string

// 		err := rows.Scan(&post_id, &user_id, &like, &dislike, &post_heading, &post_data)
// 		if err != nil {
// 			return err
// 		}
// 		fmt.Printf("%v | %v | %v | %v | %v | %v | \n", post_id, user_id, like, dislike, post_heading, post_data)
// 	}
// }
