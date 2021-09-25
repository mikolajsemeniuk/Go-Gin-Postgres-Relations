package services

import (
	"database/sql"
	"errors"
	"go-gin-postgres-relations/program/database"
	"go-gin-postgres-relations/program/inputs"
	"go-gin-postgres-relations/program/models"
)

const (
	GET_POSTS_QUERY     = "SELECT PostId, Title FROM Posts WHERE UserId = $1;"
	GET_POST_QUERY      = "SELECT PostId, Title FROM Posts WHERE PostId = $1;"
	ADD_POST_COMMAND    = "INSERT INTO Posts (UserId, Title) VALUES ($1, $2);"
	UPDATE_POST_COMMAND = "UPDATE Posts SET Title = $1 WHERE PostId = $2;"
	REMOVE_POST_COMMAND = "DELETE FROM Posts WHERE PostId = $1;"
)

func GetPostsByUserId(id int64) ([]models.Post, error) {
	var posts []models.Post

	rows, error := database.Client.Query(GET_POSTS_QUERY, id)
	if error != nil {
		return posts, error
	}
	defer rows.Close()

	for rows.Next() {
		var post models.Post
		if error := rows.Scan(&post.PostId, &post.Title); error != nil {
			return posts, error
		}
		posts = append(posts, post)
	}

	if error = rows.Err(); error != nil {
		return posts, error
	}

	return posts, nil
}

func AddPost(userId int64, input inputs.Post) error {
	if _, error := GetUser(userId); error != nil {
		return error
	}

	result, error := database.Client.Exec(ADD_POST_COMMAND, userId, input.Title)
	if error != nil {
		return error
	}

	rows, error := result.RowsAffected()
	if error != nil || rows == 0 {
		return error
	}

	return nil
}

func GetPost(id int64) (models.Post, error) {
	var post models.Post

	if error := database.Client.QueryRow(GET_POST_QUERY, id).Scan(&post.PostId, &post.Title); error != nil {
		if error == sql.ErrNoRows {
			return models.Post{}, errors.New("no record with this id :(")
		} else {
			return models.Post{}, error
		}
	}

	return post, nil
}

func UpdatePost(id int64, input inputs.Post) error {
	if _, error := GetPost(id); error != nil {
		return error
	}

	result, error := database.Client.Exec(UPDATE_POST_COMMAND, input.Title, id)
	if error != nil {
		return error
	}

	rows, error := result.RowsAffected()
	if error != nil || rows == 0 {
		return error
	}

	return nil
}

func RemovePost(id int64) error {
	if _, error := GetPost(id); error != nil {
		return error
	}

	result, error := database.Client.Exec(REMOVE_POST_COMMAND, id)
	if error != nil {
		return error
	}

	rows, error := result.RowsAffected()
	if error != nil || rows == 0 {
		return error
	}

	return nil
}
