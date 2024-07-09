package database

import (
	// "database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func InsertUser(username string, password string, email string) {
	stmt, err := db.Prepare("INSERT INTO users(username,password,email) values(?,?,?)")
	if err != nil {
		log.Fatalln(err)
	}
	_, err = stmt.Exec(username, password, email)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("Inserted User Successfully")
	}
}

func InsertCategories(category_name string) {
	stmt, err := db.Prepare("INSERT INTO categories(category_name) values(?)")
	if err != nil {
		log.Fatalln(err)
	}
	_, err = stmt.Exec(category_name)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("Inserted Categories Successfully")
	}
}

func InsertPost(user_id int, post_heading string, post_data string) {
	stmt, err := db.Prepare("INSERT INTO posts(user_id,post_heading, post_data) values (?, ?, ?)")
	if err != nil {
		log.Fatalln(err)
	}
	_, err = stmt.Exec(user_id, post_heading, post_data)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("Inserted post Successfully")
	}

	InsertPostCategories(12, 3)
}

func InsertComment(comment string, user_id int, post_id int) {
	stmt, err := db.Prepare("INSERT INTO comments(comment, user_id, post_id) values (?, ?, ?)")
	if err != nil {
		log.Fatalln(err)
	}
	_, err = stmt.Exec(comment, user_id, post_id)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("Inserted Comment Successfully")
	}
}

func InsertPostCategories(post_id int, category_id int) {
	stmt, err := db.Prepare("INSERT INTO post_categories(post_id, category_id) values (?, ?)")
	if err != nil {
		log.Fatalln(err)
	}
	_, err = stmt.Exec(post_id, category_id)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("Inserted Post with Categories Successfully")
	}
}

func InsertLikes(post_id int, user_id int) {

	stmt, err := db.Prepare("INSERT INTO likes(post_id, user_id) values (?, ?)")
	if err != nil {
		log.Fatalln(err)
	}
	_, err = stmt.Exec(post_id, user_id)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("Inserted Likes Successfully")
	}
}

func InsertDislikes(post_id int, user_id int) {
	stmt, err := db.Prepare("INSERT INTO dislikes(post_id, user_id) values (?, ?)")
	if err != nil {
		log.Fatalln(err)
	}
	_, err = stmt.Exec(post_id, user_id)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("Inserted dislikes Successfully")
	}
}

func InsertCommentDislikes(comment_id int, user_id int) {
	stmt, err := db.Prepare("INSERT INTO dislikeComment(comment_id, user_id) values (?, ?)")
	if err != nil {
		log.Fatalln(err)
	}
	_, err = stmt.Exec(comment_id, user_id)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("Inserted dislike comment Successfully")
	}
}

func InsertCommentLikes(comment_id int, user_id int) {
	stmt, err := db.Prepare("INSERT INTO likeComment(comment_id, user_id) values (?, ?)")
	if err != nil {
		log.Fatalln(err)
	}
	_, err = stmt.Exec(comment_id, user_id)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("Inserted like comment Successfully")
	}
}

func InsertSession(session string, user_id int) {
	stmt, err := db.Prepare("INSERT INTO sessions(session, user_id, timestamp) values (?, ?, ?)")
	if err != nil {
		log.Fatalln(err)
	}
	_, err = stmt.Exec(session, user_id, time.Now().Add(1*time.Minute))
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("Inserted Session Successfully")
	}
}

func AddDummyData() {

	// InsertUser(db, "zhashim", "12345678", "zahra@gmail.com")
	// InsertUser(db, "zee", "12345678", "z@gmail.com")
	// InsertUser(db, "zozo", "12345678", "zozo@gmail.com")

	// InsertCategories(db, "Sports")
	// InsertCategories(db, "Technology")
	// InsertCategories(db, "Education")
	// InsertCategories(db,"Test")

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

func ShowData() {
	rows, err := db.Query("SELECT * FROM posts", "search")
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()
	fmt.Println("DATA FOR POST")
	for rows.Next() {
		var post_id int
		var user_id int
		var dislike int
		var like int
		var post_heading string
		var post_data string

		err := rows.Scan(&post_id, &user_id, &like, &dislike, &post_heading, &post_data)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("%v | %v | %v | %v | %v | %v | \n", post_id, user_id, like, dislike, post_heading, post_data)
	}
}
