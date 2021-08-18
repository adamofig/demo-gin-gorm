package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

const (
	host     = "35.223.240.18"
	user     = ""
	password = ""
	dbname   = "test"
)

func InitDatabase() {
	if user == "" || password == "" {
		fmt.Println("RECUERDA AGREGAR LOS DATOS DE LA BASE")
	}
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=5432 sslmode=disable", host, user, password, dbname)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	fmt.Println("Database open connections")
}

func AutoMigrateModels() {
	fmt.Println("Migrando modelos si existen")
	db.AutoMigrate(&User{})
}

func GetUsersDB(c *gin.Context) {
	var users []User
	db.Find(&users)
	c.JSON(200, gin.H{
		"data": users,
	})
}

func SaveDB(c *gin.Context) {
	var user User

	c.ShouldBindJSON(&user)

	db.Create(&user)

	c.JSON(200, gin.H{
		"data": user,
	})
}

func GetUserByIdDB(c *gin.Context) {
	fmt.Println("Revisando si tengo parametros")
	id := c.Param("id")

	var user User
	// De alguna manera sabe a que tabla pertene el objeto que envio y ademas se guarda en la misma variable
	db.Find(&user, id)

	c.JSON(200, gin.H{
		"data": user,
	})

}

func DeleteDB(c *gin.Context) {
	fmt.Println("Eliminando")
	id := c.Param("id")

	var user User
	db.Delete(&user, id)

}
