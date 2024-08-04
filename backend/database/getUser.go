package database

import (
	"Forum/backend/structs"
	"log"
)

func GetUser(username string) (structs.User, error) {
	var user structs.User
	err := db.QueryRow("SELECT uid, username, email, password FROM users WHERE username = ? OR email = ?", username, username).Scan(&user.Uid, &user.Username, &user.Email, &user.Password)
	if err != nil {
		log.Println(err)
		return structs.User{}, err
	}

	return user, nil
}

func GetAllUsers() ([][]string, error) {
	rows, err := db.Query("SELECT username, email FROM users")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var users [][]string
	for rows.Next() {
		var user structs.User
		err = rows.Scan(&user.Username, &user.Email)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		users = append(users, []string{user.Username, user.Email})
	}

	if err = rows.Err(); err != nil {
		log.Println(err)
		return nil, err
	}

	return users, nil
}

func GetSession(sessionID string) (structs.Session, error) {
	var session structs.Session
	err := db.QueryRow("SELECT session_id, session, user_id, timestamp FROM sessions WHERE session = ? ", sessionID).Scan(&session.SessionID, &session.Session, &session.UserID, &session.Timestamp)
	if err != nil {
		log.Println(err)
		return structs.Session{}, err
	}
	return session, nil
}


func GetSessionByUser(userID int) (structs.Session, error) {
	var session structs.Session
	err := db.QueryRow("SELECT session_id, session, user_id, timestamp FROM sessions WHERE user_id = ? ORDER BY session_id DESC LIMIT 1 ", userID).Scan(&session.SessionID, &session.Session, &session.UserID, &session.Timestamp)
	if err != nil {
		log.Println(err)
		return structs.Session{}, err
	}
	return session, nil
}
