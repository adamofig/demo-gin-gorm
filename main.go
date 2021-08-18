package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	InitDatabase()
	AutoMigrateModels()

	router := gin.Default()

	router.GET("/", GetRootHtml)

	router.GET("/:name", IndexHandler)
	router.GET("/memory/users", GetUsers)
	router.GET("/memory/users/:id", GetUserById)

	// these methods are for CRUD to database.

	router.GET("/users", GetUsersDB)

	router.GET("/user/:id", GetUserByIdDB)
	router.POST("/user", SaveDB)
	router.DELETE("/user/:id", DeleteDB)

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
