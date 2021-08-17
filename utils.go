package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var users []User = []User{
	{ID: 1, Name: "Adamo", Company: "Data", UserType: "Provider"},
	{ID: 2, Name: "Jordan", Company: "OPtsClud", UserType: "Provider"},
	{ID: 3, Name: "Peter", Company: "OPtsClud", UserType: "Admin"},
}

func Filter(vs []User, f func(User) bool) []User {
	vsf := make([]User, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func GetUsers(c *gin.Context) {
	fmt.Println("pasando por aqui")

	c.JSON(200, gin.H{
		"data": users,
	})
}

func SaveUser(c *gin.Context) {
	name := c.Param("name")
	newUser := User{ID: 4, Name: name, Company: "new com", UserType: "provider"}
	users = append(users, newUser)

	c.JSON(200, gin.H{
		"data": users,
	})
}

func IndexHandler(c *gin.Context) {
	name := c.Param("name")

	c.JSON(200, gin.H{
		"message": name,
	})

	var data []User = nil

	print(data)
}
