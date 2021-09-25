package controllers

import (
	"go-gin-postgres-relations/program/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SetUserLike(context *gin.Context) {
	followedId, error := strconv.ParseInt(context.Param("followedid"), 10, 64)
	if error != nil {
		context.JSON(http.StatusBadRequest, "cannot convert followedid to int64")
		return
	}

	followerId, error := strconv.ParseInt(context.Param("followerid"), 10, 64)
	if error != nil {
		context.JSON(http.StatusBadRequest, "cannot convert followerid to int64")
		return
	}

	if error := services.SetUserLike(followedId, followerId); error != nil {
		context.JSON(http.StatusBadRequest, error.Error())
		return
	}

	context.JSON(http.StatusOK, "success")
}
