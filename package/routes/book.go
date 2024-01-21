package routes

import (
	"Assignment_Vivasoft/package/controllers"
	"Assignment_Vivasoft/package/middlewares"
	"net/http"
	"github.com/labstack/echo/v4"
)

type bookRoutes struct{
	echo *echo.Echo
	bookCtr controllers.BookController
	authorCtr controllers.AuthorController
	userCtr controllers.UserController
}

func BookRoutes(e *echo.Echo, bookCtr controllers.BookController, authorCtr controllers.AuthorController, userCtr controllers.UserController) *bookRoutes {
	return &bookRoutes{
		echo: e,
		bookCtr: bookCtr,
		authorCtr: authorCtr,
		userCtr: userCtr,
	}
}

func (bc *bookRoutes) InitBookRoutes(){
	e := bc.echo
	bc.initBookRoutes(e)
}

func (bc *bookRoutes) initBookRoutes(e *echo.Echo){
	// grouping route endpoints
	book := e.Group("/bookstore")

	book.GET("/ping", Pong)

	// Initializing http methods of Book - routing endpoints and their handlers
	book.POST("/book", bc.bookCtr.CreateBook, middlewares.Auth)
	book.GET("/book", bc.bookCtr.GetBooks)  // no need to authenticate
	book.PUT("/book/:bookID", bc.bookCtr.UpdateBook, middlewares.Auth)
	book.DELETE("/book/:bookID", bc.bookCtr.DeleteBook, middlewares.Auth)

	
	// Initializing http methods of Author - routing endpoints and their handlers
	book.POST("/author", bc.authorCtr.CreateAuthor)
	book.GET("/author", bc.authorCtr.GetAuthors)
	book.PUT("/author/:authorID", bc.authorCtr.UpdateAuthor)
	book.DELETE("/author/:authorID", bc.authorCtr.DeleteAuthor)


	// Initializing http methods of User - routing endpoints and their handlers
	book.POST("/user", bc.userCtr.CreateUser)
	book.GET("/user", bc.userCtr.GetUsers)
	book.PUT("/user/:userID", bc.userCtr.UpdateUser)
	book.DELETE("/user/:userID", bc.userCtr.DeleteUser)


	// Log in User
	book.POST("/login/user", bc.userCtr.LoginUser)
	
}

func Pong(e echo.Context) error{
	return e.JSON(http.StatusOK, "pong")
}