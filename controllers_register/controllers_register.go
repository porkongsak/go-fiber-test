package controllers_register

import (
	//"database/sql"
	m "go-fiber-test/models"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	//"github.com/go-playground/validator/v10"

	//"github.com/gofiber/fiber/v2/middleware/basicauth"
	"go-fiber-test/database"
	//"log"
	///"strconv"
)

func Register(c *fiber.Ctx) error {
	validate := validator.New()
	db := database.DBConn
	var user m.Users

	if err := c.BodyParser(&user); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	errors := validate.Struct(user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON("ข้อมูลผิดพลาด")
	}
	db.Create(&user)
	return c.Status(201).JSON(user)

}

func ReadUser(c *fiber.Ctx) error {
	db := database.DBConn
	search := strings.TrimSpace(c.Params("id"))
	var user []m.Users

	result := db.Find(&user, "id = ?", search)

	// returns found records count, equals `len(users)
	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(&user)
}

func DeleteUser(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var user m.Users

	result := db.Delete(&user, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}

func AllUsers(c *fiber.Ctx) error {
	db := database.DBConn
	var user []m.Users

	db.Find(&user)
	return c.Status(200).JSON(user)
}

func UpdatedUser(c *fiber.Ctx) error {
	db := database.DBConn
	var user m.Users
	id := c.Params("id")

	if err := c.BodyParser(&user); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Where("id = ?", id).Updates(&user)
	return c.Status(200).JSON(user)
}

func GetUserJson(c *fiber.Ctx) error {
	db := database.DBConn
	var users []m.Users
	// var GenZ []m.Users
	// var GenY []m.Users
	// var GenX []m.Users
	// var Baby_Boomer []m.Users
	// var G_I_Generation []m.Users

	GenZ := 0
	GenY := 0
	GenX := 0
	BabyBoomer := 0
	Generation := 0
	nerverdie := 0

	db.Find(&users)

	type UserGen struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
		Gen  string `json:"gen"`
	}
	genStr := ""
	var dataResults []UserGen

	for _, user := range users {
		if user.Age < 23 {
			genStr = "GenZ"
			GenZ += 1
			//GenZ = append(GenZ, user)
		} else if user.Age == 24 || user.Age <= 41 {
			genStr = "GenY"
			GenY += 1
			//GenY = append(GenY, user)
		} else if user.Age == 42 || user.Age <= 56 {
			genStr = "GenX"
			GenX += 1
			//GenX = append(GenX, user)
		} else if user.Age == 57 || user.Age <= 75 {
			genStr = "Baby Boomer"
			BabyBoomer += 1
			//Baby_Boomer  = append( Baby_Boomer , user)
		} else if user.Age > 76 {
			genStr = "G.I. Generation"
			Generation += 1
			//G_I_Generation  = append( G_I_Generation , user)
		} else {
			genStr = " nerver die gen"
			nerverdie += 1
		}

		d := UserGen{
			Name: user.Name,
			Age:  user.Age,
			Gen:  genStr,
		}
		dataResults = append(dataResults, d)

	}

	return c.Status(200).JSON(map[string]interface{}{
		"data":            dataResults,
		"name":            "golang-test",
		"count":           len(users), //หาผลรวม
		"GenZ":            GenZ,
		"GenY":            GenY,
		"Genx":            GenX,
		"BabyBoomer":      BabyBoomer,
		"G.I. Generation": Generation,
		"nerverdie":       nerverdie,

		//"dataGenz" GenZ,
		// "red": red,
		// "green" : green,
		// "pink" : pink,
		// "nocolor" : no,
	})

}

func SearchUsers(c *fiber.Ctx) error {

	db := database.DBConn

	search := strings.TrimSpace(c.Query("searchid"))

	var user []m.Users

	//result := db.Where("employee_id = @person OR name = @person OR lastname = @person", sql.Named("person", search)).Find(&user)

	result := db.Where("name LIKE ?", "%search%").Order(search).Find(&user)

	// returns found records count, equals `len(users)
	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(&user)
}

// func GetDogsJson(c *fiber.Ctx) error {
// 	db := database.DBConn
// 	var dogs []m.Dogs
// 	red := 0
// 	green := 0
// 	pink := 0
// 	no := 0

// 	db.Find(&dogs)
// 	// return c.Status(200).JSON(dogs)

// 	type DogsRes struct {
// 		Name  string `json:"name"`
// 		DogID int    `json:"dog_id"`
// 		Type  string `json:"type"`
// 	}

// var dataResults []DogsRes
// 	for _, v := range dogs {

// 		typeStr := ""
// 		if v.DogID <= 10 ||  v.DogID <=50 {
// 			typeStr = "red"
// 			red += 1
// 		} else if v.DogID <= 100 || v.DogID <= 150{
// 			typeStr = "green"
// 			green += 1
// 		} else if v.DogID <=200 || v.DogID <= 250 {
// 			typeStr = "pink"
// 			pink += 1
// 		} else {
// 			typeStr = "no color"
// 			no += 1
// 		}

// 		d := DogsRes{
// 			Name:  v.Name,
// 			DogID: v.DogID,
// 			Type:  typeStr,
// 		}
// 		dataResults = append(dataResults, d)
// 		// sumAmount += v.Amount
// 	}

// 	return c.Status(200).JSON(map[string]interface{}{
// 		"data":  dataResults,
// 		"name":  "golang-test",
// 		"count": len(dogs), //หาผลรวม
// 		"red": red,
// 		"green" : green,
// 		"pink" : pink,
// 		"nocolor" : no,
// 	})

// // v1.Get("/dog/json", controllers.GetDogsJson)

// }

// func AddDog(c *fiber.Ctx) error {
// 	db := database.DBConn
// 	var dog m.Dogs

// 	if err := c.BodyParser(&dog); err != nil {
// 		return c.Status(503).SendString(err.Error())
// 	}

// 	db.Create(&dog)
// 	return c.Status(201).JSON(dog)
// }
