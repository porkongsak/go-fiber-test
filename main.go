package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
   "go-fiber-test/routes"
   "go-fiber-test/database"
   m"go-fiber-test/models"

)

func initDatabase() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		"root",
		"",
		"127.0.0.1",
		"3306",
		"golang_test",
	)
	var err error
	database.DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connected!")
	database.DBConn.AutoMigrate(&m.Users {})
	database.DBConn.AutoMigrate(&m.Dogs {})
}

func main() {
	app := fiber.New()
	initDatabase()
	routes.UserRoute(app)
	app.Listen(":3000")
}






