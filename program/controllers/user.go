package controllers

import (
	"go-gin-postgres-relations/program/inputs"
	"go-gin-postgres-relations/program/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUsers(context *gin.Context) {
	users, error := services.GetUsers()
	if error != nil {
		context.JSON(http.StatusBadRequest, error.Error())
	}
	context.JSON(http.StatusOK, users)
}

func AddUser(context *gin.Context) {
	var input inputs.User

	if error := context.ShouldBindJSON(&input); error != nil {
		context.JSON(http.StatusBadRequest, "error binding data")
		return
	}

	if error := services.AddUser(input); error != nil {
		context.JSON(http.StatusBadRequest, error.Error())
		return
	}
	context.JSON(http.StatusOK, "user added")
}

func GetUser(context *gin.Context) {
	id, error := strconv.ParseInt(context.Param("id"), 10, 64)
	if error != nil {
		context.JSON(http.StatusBadRequest, "cannot convert id to int64")
		return
	}

	user, error := services.GetUser(id)
	if error != nil {
		context.JSON(http.StatusBadRequest, error.Error())
		return
	}

	context.JSON(http.StatusOK, user)
}

func UpdateUser(context *gin.Context) {
	id, error := strconv.ParseInt(context.Param("id"), 10, 64)
	if error != nil {
		context.JSON(http.StatusBadRequest, "cannot convert id to int64")
		return
	}

	var input inputs.User
	if error := context.ShouldBindJSON(&input); error != nil {
		context.JSON(http.StatusBadRequest, "error binding data")
		return
	}

	if error := services.UpdateUser(id, input); error != nil {
		context.JSON(http.StatusBadRequest, error.Error())
		return
	}

	context.JSON(http.StatusOK, "user updated")
}

func RemoveUser(context *gin.Context) {
	id, error := strconv.ParseInt(context.Param("id"), 10, 64)
	if error != nil {
		context.JSON(http.StatusBadRequest, "cannot convert id to int64")
		return
	}

	if error := services.RemoveUser(id); error != nil {
		context.JSON(http.StatusBadRequest, error.Error())
		return
	}

	context.JSON(http.StatusOK, "user removed")
}
