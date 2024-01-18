package controllers

import (
	"Assignment_Vivasoft/package/domain"
	"Assignment_Vivasoft/package/models"
	"Assignment_Vivasoft/package/types"
	"net/http"
	"strconv"
	"github.com/labstack/echo/v4"
)

type IBookController interface {
	GetBooks(e echo.Context) error
	CreateBook(e echo.Context) error
	UpdateBook(e echo.Context) error
	DeleteBook(e echo.Context) error
}

type BookController struct {
	bookSvc domain.IBookService
}

func NewBookController(bookSvc domain.IBookService) BookController {
	return BookController{
		bookSvc: bookSvc,
	}
}


// CreateBook implements IBookController.
func (controller *BookController) CreateBook(e echo.Context) error {
	reqBook := &types.BookRequest{}
	if err := e.Bind(reqBook); err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid Data")
	}
	if err := reqBook.Validate(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	book := &models.BookDetails{
		BookName:    reqBook.BookName,
		AuthorId:    reqBook.AuthorId,
		Publication: reqBook.Publication,
	}
	if err := controller.bookSvc.CreateBook(book); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusCreated, "BookDetail is created successfully")
}

// GetBooks implements IBookController.
func (controller *BookController) GetBooks(e echo.Context) error {
	tempBookID := e.QueryParam("bookID")
	bookID, err := strconv.ParseInt(tempBookID, 0, 0)
	if err != nil && tempBookID != "" {
		return e.JSON(http.StatusBadRequest, "Enter a valid book ID")
	}
	book, err := controller.bookSvc.GetBooks(uint(bookID))
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, book)
}

// UpdateBook implements IBookController.
func (controller *BookController) UpdateBook(e echo.Context) error {
	reqBook := &types.BookRequest{}
	if err := e.Bind(reqBook); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	if err := reqBook.Validate(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	tempBookID := e.Param("bookID")
	bookID, err := strconv.ParseInt(tempBookID, 0, 0)
	if err != nil {
		return e.JSON(http.StatusBadRequest, "Enter a valid book ID")
	}
	existingBook, err := controller.bookSvc.GetBooks(uint(bookID))
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	updateBook := &models.BookDetails{
		BookId:   uint(bookID),
		BookName: reqBook.BookName,
		AuthorId: reqBook.AuthorId,
		Publication: reqBook.Publication,
	}

	if updateBook.BookName == "" {
		updateBook.BookName = existingBook[0].BookName
	}
	if updateBook.AuthorId == 0 {
		updateBook.AuthorId = existingBook[0].AuthorId
	}
	if updateBook.Publication == "" {
		updateBook.Publication = existingBook[0].Publication
	}

	if err := controller.bookSvc.UpdateBook(updateBook); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusCreated, "BookDetail is updated successfully")
}

// DeleteBook implements IBookController.
func (controller *BookController) DeleteBook(e echo.Context) error {
	tempBookID := e.Param("bookID")
	bookID, err := strconv.ParseInt(tempBookID, 0, 0)
	if err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid Data")
	}	
	_, err = controller.bookSvc.GetBooks(uint(bookID))
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	if err := controller.bookSvc.DeleteBook(uint(bookID)); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, "BookDetail is deleted successfully")
}

// DeleteBook using Author ID
func (controller *BookController) DeleteBookByAuthorID(e echo.Context) error {
	tempAuthorID := e.Param("authorID")
	authorID, err := strconv.ParseInt(tempAuthorID, 0, 0)
	if err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid Data")
	}	
	_, err = controller.bookSvc.GetBooks(uint(authorID))
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	if err := controller.bookSvc.DeleteBook(uint(authorID)); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, "BookDetail is deleted successfully")
}
