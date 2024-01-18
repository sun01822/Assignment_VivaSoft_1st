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
	config.InitConfig()

	// Database initializations
	db := connection.GetDB()
	if db == nil{
		panic("Database connection failed")
	}
	fmt.Println("Database connected successfully")

	// Repository initializations
	bookRepository := repositories.BookDBInstance(db)
	authorRepository := repositories.AuthorDBInstance(db)

	// Service initializations
	bookService := services.BookServiceInstance(bookRepository)
	authorService := services.AuthorServiceInstance(authorRepository)

	// Controllers initialization
	bookController := controllers.NewBookController(bookService)
	authorController := controllers.NewAuthorController(authorService, bookService)

	// Route initializations
	book := routes.BookRoutes(e, bookController, authorController)


	book.InitBookRoutes()
	// Starting Server 
	log.Fatal(e.Start(fmt.Sprintf(":%s", config.LocalConfig.Port)))

}