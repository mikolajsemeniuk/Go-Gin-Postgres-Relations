package services

import (
	"errors"
	"fmt"
	"go-gin-postgres-relations/program/database"
)

const (
	CHECK_USERLIKE_QUERY    = "SELECT COUNT(*) FROM UserLikes WHERE FollowedId = $1 AND FollowerId = $2;"
	ADD_USERLIKE_COMMAND    = "INSERT INTO UserLikes (FollowedId, FollowerId) VALUES ($1, $2);"
	REMOVE_USERLIKE_COMMAND = "DELETE FROM UserLikes WHERE FollowedId = $1 AND FollowerId = $2;"
)

func CheckIfUserLikeExists(followedId int64, followerId int64) (bool, error) {
	var ifExists bool

	if error := database.Client.QueryRow(CHECK_USERLIKE_QUERY, followedId, followerId).Scan(&ifExists); error != nil {
		return ifExists, error
	}

	fmt.Println("ifExists", ifExists, ", followedId: ", followedId, ", followerId: ", followerId)

	return ifExists, nil
}

func SetUserLike(followedId int64, followerId int64) error {
	if followedId == followerId {
		return errors.New("you can't like yourself")
	}

	if _, error := GetUser(followedId); error != nil {
		return error
	}

	if _, error := GetUser(followerId); error != nil {
		return error
	}

	exists, error := CheckIfUserLikeExists(followedId, followerId)
	if error != nil {
		return error
	}

	var command string
	if exists {
		command = REMOVE_USERLIKE_COMMAND
	} else {
		command = ADD_USERLIKE_COMMAND
	}

	result, error := database.Client.Exec(command, followedId, followerId)
	if error != nil {
		return error
	}

	rows, error := result.RowsAffected()
	if error != nil || rows == 0 {
		return error
	}

	return nil
}
