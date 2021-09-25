package services

import (
	"errors"
	"go-gin-postgres-relations/program/database"
	"go-gin-postgres-relations/program/payloads"
)

const (
	GET_POSTLIKES_QUERY     = "SELECT UserName FROM PostLikes JOIN Users ON PostLikes.UserId = Users.UserId WHERE PostLikes.PostId = $1;"
	GET_POSTLIKE_QUERY      = "SELECT COUNT(*) FROM PostLikes WHERE UserId = $1 AND PostId = $2;"
	CHECK_POSTLIKE_QUERY    = "SELECT COUNT(*) FROM Posts WHERE UserId = $1 AND PostId = $2;"
	ADD_POSTLIKE_COMMAND    = "INSERT INTO PostLikes (UserId, PostId) VALUES ($1, $2);"
	REMOVE_POSTLIKE_COMMAND = "DELETE FROM PostLikes WHERE UserId = $1 AND PostId = $2;"
)

func CheckPostLikeExists(userid int64, postid int64) (bool, error) {
	var ifExists bool

	if error := database.Client.QueryRow(GET_POSTLIKE_QUERY, userid, postid).Scan(&ifExists); error != nil {
		return ifExists, error
	}

	return ifExists, nil
}

func CheckIfUserLikesHisOwnPost(userid int64, postid int64) (bool, error) {
	var ifExists bool

	if error := database.Client.QueryRow(CHECK_POSTLIKE_QUERY, userid, postid).Scan(&ifExists); error != nil {
		return ifExists, error
	}

	return ifExists, nil
}

func SetPostLike(userid int64, postid int64) error {

	if _, error := GetUser(userid); error != nil {
		return error
	}

	if _, error := GetPost(postid); error != nil {
		return error
	}

	exists, error := CheckIfUserLikesHisOwnPost(userid, postid)
	if error != nil {
		return error
	}

	if exists {
		return errors.New("you cannot like your own posts")
	}

	exists, error = CheckPostLikeExists(userid, postid)
	if error != nil {
		return error
	}

	var command string
	if exists {
		command = REMOVE_POSTLIKE_COMMAND
	} else {
		command = ADD_POSTLIKE_COMMAND
	}

	result, error := database.Client.Exec(command, userid, postid)
	if error != nil {
		return error
	}

	rows, error := result.RowsAffected()
	if error != nil || rows == 0 {
		return error
	}

	return nil
}

func GetPostLikesByPostId(postid int64) ([]payloads.PostLike, error) {
	var postlikes []payloads.PostLike

	rows, error := database.Client.Query(GET_POSTLIKES_QUERY, postid)
	if error != nil {
		return postlikes, error
	}
	defer rows.Close()

	for rows.Next() {
		var postlike payloads.PostLike
		if error := rows.Scan(&postlike.Username); error != nil {
			return postlikes, error
		}
		postlikes = append(postlikes, postlike)
	}

	if error = rows.Err(); error != nil {
		return postlikes, error
	}

	return postlikes, nil
}
