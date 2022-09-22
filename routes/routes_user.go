package routes

import (
	"github.com/gofiber/fiber/v2"

	"go-fiber-test/controllers"
	"go-fiber-test/controllers_register"

	//"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func UserRoute(app *fiber.App) {

	api := app.Group("/api") // /api    api/v2/
	
	v2 := api.Group("/v2")
	v2.Post("/inet", controllers.Query)
	v2.Get("/all", controllers_register.AllUsers)  // all user
	v2.Get("/read-user/:id", controllers_register.ReadUser) // read id users
	v2.Get("/user/json", controllers_register.GetUserJson)
	v2.Get("/user/filter", controllers_register.SearchUsers)
	


// Provide a minimal config
app.Use(basicauth.New(basicauth.Config{
	Users: map[string]string{
		"john":  "doe", //1
		"admin": "123456",  //2
		"testgo": "2192565",
	},
}))

 // /api/v2
 v1 := api.Group("/v1") // /api/v1
	// api/v1
v1.Get("/",controllers.Helo )
//  ส่งผ่าน json    
v1.Post("/",controllers.Bodypasder )
v1.Post("/regis", controllers.Register)

// CRUD User
v1.Post("/register-user", controllers_register.Register)
v1.Delete("/delete-user/:id", controllers_register.DeleteUser)
v1.Put("/update-user/:id", controllers_register.UpdatedUser)



v1.Get("/dog/json", controllers.GetDogsJson) ///

// รับผ่าน พนีะำพ
v1.Get("/user/:name", controllers.Params )
v1.Get("/fact/:id", controllers.ParamsFac )

// c.query  Query Params  api/v2
v1.Post("/inet", controllers.QueryV2)


v1.Get("/dog", controllers.GetDogs)
	v1.Get("/dog/filter", controllers.GetDog)
	v1.Post("/dog", controllers.AddDog)
	v1.Put("/dog/:id", controllers.UpdateDog)
	v1.Delete("/dog/:id", controllers.RemoveDog)


//v2.Get("/user/:number", controllers.ParamsFac )


  

// post 

// //Post addข้อมูล 
// //Put updateข้อมูล
// //Delte deleteข้อมูล


//  v2.Put("/user",)
//  v2.Delete("/user",)

}