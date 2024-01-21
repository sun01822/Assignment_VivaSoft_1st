package containers

import (
	"Assignment_Vivasoft/package/config"
	"Assignment_Vivasoft/package/connection"
	"Assignment_Vivasoft/package/controllers"
	"Assignment_Vivasoft/package/repositories"
	"Assignment_Vivasoft/package/routes"
	"Assignment_Vivasoft/package/services"
	"fmt"
	"log"
	"github.com/labstack/echo/v4"
)

func Serve(e *echo.Echo){
	
	// config initialization
	config.SetConfig()

	// Database initializations
	db := connection.GetDB()

	// Repository initializations
	bookRepository := repositories.BookDBInstance(db)
	authorRepository := repositories.AuthorDBInstance(db)
	userRepository := repositories.UserDBInstance(db)

	// Service initializations
	bookService := services.BookServiceInstance(bookRepository)
	authorService := services.AuthorServiceInstance(authorRepository)
	userService := services.UserServiceInstance(userRepository)


	// Controllers initialization
	bookController := controllers.NewBookController(bookService)
	authorController := controllers.NewAuthorController(authorService, bookService)
	userController := controllers.NewUserController(userService)

	// Route initializations
	book := routes.BookRoutes(e, bookController, authorController, userController)


	book.InitBookRoutes()
	// Starting Server 
	log.Fatal(e.Start(fmt.Sprintf(":%s", config.LocalConfig.Port)))

}