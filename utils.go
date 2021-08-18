package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var users []User = []User{
	{Name: "Adamo", Company: "GNP", UserType: "Provider"},
	{Name: "Jordan", Company: "CBRE", UserType: "Provider"},
	{Name: "Peter", Company: "Ksquare", UserType: "Admin"},
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

	c.JSON(200, gin.H{
		"data": users,
	})
}

func GetUserById(c *gin.Context) {
	fmt.Println(c.Param("id"))

	user := Filter(users, func(u User) bool {
		fmt.Println("user", u, "Id is", u.ID, fmt.Sprint(u.ID))

		if fmt.Sprint(u.ID) == c.Param("id") {
			return true
		} else {
			return false
		}
	})

	c.JSON(200, gin.H{
		"data": user,
	})
}

func SaveUser(c *gin.Context) {
	name := c.Param("name")
	newUser := User{Name: name, Company: "new com", UserType: "provider"}
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

func GetRootHtml(router *gin.Context) {
	html := `
	<h3> Bienvenido a la API CRUD  </h3>
	<h4> Get Users  <a href="/memory/users"> /memory/users </a> </h4>
	<h4> Configuraciones .... /config</h4>
	`

	router.Data(200, "text/html; charset=utf-8", []byte(html))
}
