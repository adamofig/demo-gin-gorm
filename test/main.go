package main

import (
	"fmt"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Company  string `json:"company"`
	UserType string `json:"userType"`
}

var users []User = []User{
	{ID: 1, Name: "Adamo", Company: "Data", UserType: "Provider"},
	{ID: 2, Name: "Jordan", Company: "OPtsClud", UserType: "Provider"},
	{ID: 3, Name: "Peter", Company: "OPtsClud", UserType: "Admin"},
}

func Filter(users []User, f func(User) bool) []User {
	vsf := make([]User, 0)
	for _, v := range users {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func main() {
	fmt.Println(users)

	filtered := Filter(users, func(u User) bool { return u.Name == "Adamo" })

	fmt.Println(filtered)

	// TODO agregar la cadena de conexion aqui.
	dsn := ""
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	fmt.Println(db, err)

	SaveUser(db)

	// Usé esta instrucción para migrar la tabla a la base de datos.
	// db.AutoMigrate(&UserDB{})

}

func GetUserDB(db *gorm.DB) {
	var user []User
	db.Find(&user) // todavia no entiendo como esto puede referenciar a la base
}

func SaveUser(db *gorm.DB) {
	user := UserDB{Name: "pedro"}
	db.Create(&user)
}
