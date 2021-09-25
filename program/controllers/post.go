package controllers

import (
	"fmt"
	"go-gin-postgres-relations/program/inputs"
	"go-gin-postgres-relations/program/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddPost(context *gin.Context) {
	id, error := strconv.ParseInt(context.Param("id"), 10, 64)
	if error != nil {
		context.JSON(http.StatusBadRequest, "cannot convert id to int64")
		return
	}

	var input inputs.Post
	if error := context.ShouldBindJSON(&input); error != nil {
		context.JSON(http.StatusBadRequest, "error binding data")
		return
	}

	if error := services.AddPost(id, input); error != nil {
		context.JSON(http.StatusBadRequest, error.Error())
		return
	}
	context.JSON(http.StatusOK, fmt.Sprintf("post added to user of id: %d", id))
}

func GetPost(context *gin.Context) {
	id, error := strconv.ParseInt(context.Param("id"), 10, 64)
	if error != nil {
		context.JSON(http.StatusBadRequest, "cannot convert id to int64")
		return
	}

	post, error := services.GetPost(id)
	if error != nil {
		context.JSON(http.StatusBadRequest, error.Error())
		return
	}

	context.JSON(http.StatusOK, post)
}

func UpdatePost(context *gin.Context) {
	id, error := strconv.ParseInt(context.Param("id"), 10, 64)
	if error != nil {
		context.JSON(http.StatusBadRequest, "cannot convert id to int64")
		return
	}

	var input inputs.Post
	if error := context.ShouldBindJSON(&input); error != nil {
		context.JSON(http.StatusBadRequest, "error binding data")
		return
	}

	if error := services.UpdatePost(id, input); error != nil {
		context.JSON(http.StatusBadRequest, error.Error())
		return
	}

	context.JSON(http.StatusOK, "post updated")
}

func RemovePost(context *gin.Context) {
	id, error := strconv.ParseInt(context.Param("id"), 10, 64)
	if error != nil {
		context.JSON(http.StatusBadRequest, "cannot convert id to int64")
		return
	}

	if error := services.RemovePost(id); error != nil {
		context.JSON(http.StatusBadRequest, error.Error())
		return
	}

	context.JSON(http.StatusOK, "post removed")
}
