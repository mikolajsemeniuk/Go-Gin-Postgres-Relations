package application

import (
	"go-gin-postgres-relations/program/controllers"
	_ "go-gin-postgres-relations/program/docs" // swag init generates this folder

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	DOCS = "http://localhost:3000/swagger/doc.json"
)

func Route() {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL(DOCS)))

	// USER
	router.GET("/user", controllers.GetUsers)
	router.POST("/user", controllers.AddUser)
	router.GET("/user/:id", controllers.GetUser)
	router.PATCH("/user/:id", controllers.UpdateUser)
	router.DELETE("/user/:id", controllers.RemoveUser)

	// POST
	// NORMALLY FOR POST `id` goes from JWT but I didn't implement it, it goes from route
	router.POST("/post/:id", controllers.AddPost)
	router.GET("/post/:id", controllers.GetPost)
	router.PATCH("/post/:id", controllers.UpdatePost)
	router.DELETE("/post/:id", controllers.RemovePost)

	// USERLIKES
	router.PATCH("/userlike/:followedid/:followerid", controllers.SetUserLike)

	// POSTLIKES
	router.PATCH("/postlike/:userid/:postid", controllers.SetPostLike)
}
