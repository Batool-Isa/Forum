package Forum

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InsertUser(db *sql.DB, username string, password string, email string) {
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

func InsertCategories(db *sql.DB, category_name string) {
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

func InsertPost(db *sql.DB, user_id int, post_heading string, post_data string) {
	stmt, err := db.Prepare("INSERT INTO posts(user_id,post_heading, post_data) values (?, ?, ?)")
	if err != nil {
		log.Fatalln(err)
	}
	_, err = stmt.Exec(user_id, post_heading, post_data)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("Inserted Categories Successfully")
	}
}

func InsertComment(db *sql.DB, comment string, user_id int, post_id int) {
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

func InsertPostCategories(db *sql.DB, post_id int, category_id int) {
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

func AddDummyData(db *sql.DB) {

	// InsertUser(db, "zhashim", "12345678", "zahra@gmail.com")
	// InsertUser(db, "zee", "12345678", "z@gmail.com")

	// InsertCategories(db, "Sports")
	// InsertCategories(db, "Technology")
	// InsertCategories(db, "Education")

	// InsertPost(db, 1, "Advancements in AI Technology", "Explore the latest advancements in artificial intelligence and its applications.")
	// InsertPost(db, 1, "Online Learning Platforms", "Discover the best online learning platforms to enhance your skills and knowledge.")

	// InsertPostCategories(db,2,3)
	// InsertPostCategories(db,1,2)

}

func ShowData(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM posts", "search")
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()

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
		fmt.Printf("User %d performed search: %v\n", user_id,post_id)
	}
}