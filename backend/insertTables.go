package Forum

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InsertUser(db *sql.DB, username string, password string, email string) {
	stmt, err := db.Prepare("INSERT INTO user(username,password,email) values(?,?,?)")
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

func InsertPost(db *sql.DB, uid int, user_id int, post_data string) {
	stmt, err := db.Prepare("INSERT INTO post(uid, user_id, post_data) values (?, ?, ?)")
	if err != nil {
		log.Fatalln(err)
	}
	_, err = stmt.Exec(uid, user_id, post_data)
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
		log.Println("Inserted Categories Successfully")
	}
}

func AddDummyData(db *sql.DB) {
	
}
