package services

import (
	"database/sql"
	"errors"
	"go-gin-postgres-relations/program/database"
	"go-gin-postgres-relations/program/inputs"
	"go-gin-postgres-relations/program/models"
)

const (
	GET_USERS_QUERY     = "SELECT UserId, Username FROM Users;"
	GET_USER_QUERY      = "SELECT UserId, Username FROM Users WHERE UserId = $1;"
	ADD_USER_COMMAND    = "INSERT INTO Users (Username) VALUES ($1);"
	UPDATE_USER_COMMAND = "UPDATE Users SET Username = $1 WHERE UserId = $2;"
	REMOVE_USER_COMMAND = "DELETE FROM Users WHERE UserId = $1;"
)

func GetUsers() ([]models.User, error) {
	var users []models.User

	rows, error := database.Client.Query(GET_USERS_QUERY)
	if error != nil {
		return users, error
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User

		if error := rows.Scan(&user.UserId, &user.Username); error != nil {
			return users, error
		}

		posts, error := GetPostsByUserId(user.UserId)
		if error != nil {
			return users, error
		}

		if len(posts) == 0 {
			user.Posts = []models.Post{}
		} else {
			user.Posts = posts
		}

		users = append(users, user)
	}

	if error = rows.Err(); error != nil {
		return users, error
	}

	return users, nil
}

func AddUser(input inputs.User) error {
	result, error := database.Client.Exec(ADD_USER_COMMAND, input.Username)
	if error != nil {
		return error
	}

	rows, error := result.RowsAffected()
	if error != nil || rows == 0 {
		return error
	}

	return nil
}

func GetUser(id int64) (models.User, error) {
	var user models.User

	if error := database.Client.QueryRow(GET_USER_QUERY, id).Scan(&user.UserId, &user.Username); error != nil {
		if error == sql.ErrNoRows {
			return models.User{}, errors.New("no record with this id :(")
		} else {
			return models.User{}, error
		}
	}

	posts, error := GetPostsByUserId(user.UserId)
	if error != nil {
		return user, error
	}

	if len(posts) == 0 {
		user.Posts = []models.Post{}
	} else {
		user.Posts = posts
	}

	return user, nil
}

func UpdateUser(id int64, input inputs.User) error {
	if _, error := GetUser(id); error != nil {
		return error
	}

	result, error := database.Client.Exec(UPDATE_USER_COMMAND, input.Username, id)
	if error != nil {
		return error
	}

	rows, error := result.RowsAffected()
	if error != nil || rows == 0 {
		return error
	}
	return nil
}

func RemoveUser(id int64) error {
	if _, error := GetUser(id); error != nil {
		return error
	}

	result, error := database.Client.Exec(REMOVE_USER_COMMAND, id)
	if error != nil {
		return error
	}

	rows, error := result.RowsAffected()
	if error != nil || rows == 0 {
		return error
	}
	return nil
}
