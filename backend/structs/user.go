package structs

import "time"

type User struct {
	Uid			int
	Username 	string
	Password 	string
	Email		string
}


type Session struct {
	SessionID    int
	Session		 string
	Timestamp    time.Time
	UserID       int
}