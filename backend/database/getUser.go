package database

import ("Forum/backend/structs"
"fmt")
func GetUser(username string) (structs.User, error) {
	var user structs.User
    err := db.QueryRow("SELECT uid, username, email, password FROM users WHERE username = ?", username).Scan(&user.Uid, &user.Username, &user.Email, &user.Password)
    fmt.Println(err)
    if err != nil {
        return structs.User{}, err
	}

	return user, nil
}

func GetAllUsers() ([]string, error) {
	rows, err := db.Query("SELECT username FROM users")
	if err != nil {
        return nil, err
	}
	defer rows.Close()

	var users []string
	if rows.Next() {
        var user structs.User
        err = rows.Scan(&user.Username)
        if err != nil {
            return nil, err
        }
        users = append(users, user.Username)
    } 

    if err = rows.Err(); err != nil {
		return nil, err
	}


	return users, nil
}
