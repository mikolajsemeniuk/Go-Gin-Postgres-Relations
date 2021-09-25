package services

import (
	"database/sql"
	"fmt"
	"go-gin-postgres-relations/program/database"
	"go-gin-postgres-relations/program/inputs"
	"go-gin-postgres-relations/program/payloads"
)

const (
	GET_USERS_QUERY     = "SELECT UserId, Username FROM Users;"
	GET_USER_QUERY      = "SELECT UserId, Username FROM Users WHERE UserId = $1;"
	GET_FOLLOWED_QUERY  = "SELECT UserName from USERS JOIN UserLikes on Users.UserId = FollowerId WHERE FollowedId = $1;"
	GET_FOLLOWERS_QUERY = "SELECT UserName from USERS JOIN UserLikes on Users.UserId = FollowedId WHERE FollowerId = $1;"
	ADD_USER_COMMAND    = "INSERT INTO Users (Username) VALUES ($1);"
	UPDATE_USER_COMMAND = "UPDATE Users SET Username = $1 WHERE UserId = $2;"
	REMOVE_USER_COMMAND = "DELETE FROM Users WHERE UserId = $1;"
)

func GetUsers() ([]payloads.User, error) {
	var users []payloads.User

	rows, error := database.Client.Query(GET_USERS_QUERY)
	if error != nil {
		return users, error
	}
	defer rows.Close()

	for rows.Next() {
		var user payloads.User

		if error := rows.Scan(&user.UserId, &user.Username); error != nil {
			return users, error
		}

		posts, error := GetPostsByUserId(user.UserId)
		if error != nil {
			return users, error
		}

		if len(posts) == 0 {
			user.Posts = []payloads.Post{}
		} else {
			user.Posts = posts
		}

		followeds, error := GetFollowers(user.UserId, GET_FOLLOWED_QUERY)
		if error != nil {
			return users, error
		}

		if len(followeds) == 0 {
			user.Followed = []payloads.Follower{}
		} else {
			user.Followed = followeds
		}

		followers, error := GetFollowers(user.UserId, GET_FOLLOWERS_QUERY)
		if error != nil {
			return users, error
		}

		if len(followers) == 0 {
			user.Followers = []payloads.Follower{}
		} else {
			user.Followers = followers
		}

		users = append(users, user)
	}

	if error = rows.Err(); error != nil {
		return users, error
	}

	return users, nil
}

func GetFollowers(id int64, query string) ([]payloads.Follower, error) {
	var followeds []payloads.Follower

	rows, error := database.Client.Query(query, id)
	if error != nil {
		return followeds, error
	}
	defer rows.Close()

	for rows.Next() {
		var followed payloads.Follower

		if error := rows.Scan(&followed.Username); error != nil {
			return followeds, error
		}

		followeds = append(followeds, followed)
	}

	if error = rows.Err(); error != nil {
		return followeds, error
	}

	return followeds, nil
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

func GetUser(id int64) (payloads.User, error) {
	var user payloads.User

	if error := database.Client.QueryRow(GET_USER_QUERY, id).Scan(&user.UserId, &user.Username); error != nil {
		if error == sql.ErrNoRows {
			return payloads.User{}, fmt.Errorf("no user with id: %d", id)
		} else {
			return payloads.User{}, error
		}
	}

	posts, error := GetPostsByUserId(user.UserId)
	if error != nil {
		return user, error
	}

	if len(posts) == 0 {
		user.Posts = []payloads.Post{}
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
