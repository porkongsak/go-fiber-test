package controllers

import (
	m "go-fiber-test/models"
	"strings"
	
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	
	//"github.com/go-playground/validator/v10"

	//"github.com/gofiber/fiber/v2/middleware/basicauth"
	"go-fiber-test/database"
	"log"
	"strconv"
)

func Helo(c *fiber.Ctx) error {
	return c.SendString("rf, World!")
}

func Bodypasder(c *fiber.Ctx) error {

	p := new(m.Person)

	if err := c.BodyParser(p); err != nil {
		return err
	}

	log.Println(p.Name) // john
	log.Println(p.Pass) // doe

	return c.JSON(p)
}

func Params(c *fiber.Ctx) error {
	c.Params("name") // "fenny"
	a := c.Params("name")
	str := a
	return c.JSON(str)

}

// //
func ParamsFac(c *fiber.Ctx) error {
	c.Params("id") // "fenny"
	a := c.Params("id")
	num, _ := strconv.Atoi(a)
	str := factorial(num)
	return c.JSON(str)

}

func factorial(num int) int {
	if num == 1 || num == 0 {
		return num
	}
	return num * factorial(num-1)
}

func Query(c *fiber.Ctx) error {
	c.Query("search") // "fenny"
	a := c.Query("search")
	str := " => v2 hello my name is" + a
	return c.JSON(str)

}

func QueryV2(c *fiber.Ctx) error {
	c.Query("search") // "fenny"
	a := c.Query("search")
	str := " => v2 hello my name is" + a
	return c.JSON(str)

}

func Register(c *fiber.Ctx) error {
	//Connect to database
	//db := database.DBConn
	validate := validator.New()

	var user m.Users

	if err := c.BodyParser(&user); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	//db.Create(&user)


	errors := validate.Struct(user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON("ข้อมูลผิดพลาด")
	}
	//db.Create(&user)
	return c.Status(201).JSON(user)

	// p := new(m.Users)

	// if err := c.BodyParser(p); err != nil {
	// 	return err
	// }

	// log.Println(p.Name) // john
	// log.Println(p.Pass) // doe

	// return c.JSON(p)
}

func GetDogs(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Find(&dogs)
	return c.Status(200).JSON(dogs)
}


// Query  postman params
func GetDog(c *fiber.Ctx) error {
	db := database.DBConn
	search := strings.TrimSpace(c.Query("search"))
	var dog []m.Dogs

	result := db.Find(&dog, "dog_id = ?", search)

	// returns found records count, equals `len(users)
	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(&dog)
}




// /////////////
func AddDog(c *fiber.Ctx) error {
	db := database.DBConn
	var dog m.Dogs

	if err := c.BodyParser(&dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Create(&dog)
	return c.Status(201).JSON(dog)
}

func UpdateDog(c *fiber.Ctx) error {
	db := database.DBConn
	var dog m.Dogs
	id := c.Params("id")

	if err := c.BodyParser(&dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Where("id = ?", id).Updates(&dog)
	return c.Status(200).JSON(dog)
}

func RemoveDog(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var dog m.Dogs

	result := db.Delete(&dog, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}
   



func GetDogsJson(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs
	red := 0
	green := 0
	pink := 0
	no := 0
	

	db.Find(&dogs)
	// return c.Status(200).JSON(dogs)

	type DogsRes struct {
		Name  string `json:"name"`
		DogID int    `json:"dog_id"`
		Type  string `json:"type"`
	}

var dataResults []DogsRes
	for _, v := range dogs {

		typeStr := ""
		if v.DogID <= 10 ||  v.DogID <=50 {
			typeStr = "red"
			red += 1
		} else if v.DogID <= 100 || v.DogID <= 150{
			typeStr = "green"
			green += 1
		} else if v.DogID <=200 || v.DogID <= 250 {
			typeStr = "pink"
			pink += 1
		} else {
			typeStr = "no color"
			no += 1
		}

		d := DogsRes{
			Name:  v.Name,
			DogID: v.DogID,
			Type:  typeStr,
		}
		dataResults = append(dataResults, d)
		// sumAmount += v.Amount
	}

	return c.Status(200).JSON(map[string]interface{}{
		"data":  dataResults,
		"name":  "golang-test",
		"count": len(dogs), //หาผลรวม
		"red": red,
		"green" : green,
		"pink" : pink,
		"nocolor" : no,
	})


// v1.Get("/dog/json", controllers.GetDogsJson)

}


// var dataResults []DogsRes
// 	for _, v := range dogs {

// 		typeStr := ""
// 		if v.DogID == 111 {
// 			typeStr = "red"
// 		} else if v.DogID == 113 {
// 			typeStr = "green"
// 		} else if v.DogID == 999 {
// 			typeStr = "pink"
// 		} else {
// 			typeStr = "no color"
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
// "count": len(dogs), //หาผลรวม
// 	})


// v1.Get("/dog/json", controllers.GetDogsJson)
