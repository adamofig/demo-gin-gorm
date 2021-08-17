package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/:name", IndexHandler)
	router.GET("/user", GetUsers)
	router.GET("/newUser/:name", SaveUser)
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Company  string `json:"company"`
	UserType string `json:"userType"`
}
