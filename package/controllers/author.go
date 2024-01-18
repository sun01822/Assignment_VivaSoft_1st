package controllers

import (
	"Assignment_Vivasoft/package/domain"
	"Assignment_Vivasoft/package/models"
	"Assignment_Vivasoft/package/types"
	"net/http"
	"strconv"
	"github.com/labstack/echo/v4"
)

type IAuthorController interface {
	GetAuthors(e echo.Context) error
	CreateAuthor(e echo.Context) error
	UpdateAuthor(e echo.Context) error
	DeleteAuthor(e echo.Context) error
	DeleteBookByAuthorID(e echo.Context) error
}

type AuthorController struct {
	AuthorSvc domain.IAuthorService
	BookSvc   domain.IBookService // Add the missing BookSvc field
}

func NewAuthorController(AuthorSvc domain.IAuthorService, BookSvc domain.IBookService) AuthorController {
	return AuthorController{
		AuthorSvc: AuthorSvc,
		BookSvc:   BookSvc, // Initialize the BookSvc field
	}
}


// CreateAuthor implements IAuthorController.
func (controller *AuthorController) CreateAuthor(e echo.Context) error {
	reqAuthor := &types.AuthorRequest{}
	if err := e.Bind(reqAuthor); err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid Data")
	}
	if err := reqAuthor.Validate(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	Author := &models.AuthorDetails{
		 AuthorName: reqAuthor.AuthorName,
		 AuthorAddress: reqAuthor.AuthorAddress,
		 AuthorPhone: reqAuthor.AuthorPhone,
	}
	if err := controller.AuthorSvc.CreateAuthor(Author); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusCreated, "AuthorDetail is created successfully")
}

// GetAuthors implements IAuthorController.
func (controller *AuthorController) GetAuthors(e echo.Context) error {
	tempAuthorID := e.QueryParam("authorID")
	AuthorID, err := strconv.ParseInt(tempAuthorID, 0, 0)
	if err != nil && tempAuthorID != "" {
		return e.JSON(http.StatusBadRequest, "Enter a valid Author ID")
	}
	Author, err := controller.AuthorSvc.GetAuthors(uint(AuthorID))
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, Author)
}

// UpdateAuthor implements IAuthorController.
func (controller *AuthorController) UpdateAuthor(e echo.Context) error {
	reqAuthor := &types.AuthorRequest{}
	if err := e.Bind(reqAuthor); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	if err := reqAuthor.Validate(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	tempAuthorID := e.Param("authorID")
	AuthorID, err := strconv.ParseInt(tempAuthorID, 0, 0)
	if err != nil {
		return e.JSON(http.StatusBadRequest, "Enter a valid Author ID")
	}
	existingAuthor, err := controller.AuthorSvc.GetAuthors(uint(AuthorID))
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	updateAuthor := &models.AuthorDetails{
		AuthorId:    uint(AuthorID),
		AuthorName:   reqAuthor.AuthorName,
		AuthorAddress: reqAuthor.AuthorAddress,
		AuthorPhone:  reqAuthor.AuthorPhone,
	}

	if updateAuthor.AuthorName == "" {
		updateAuthor.AuthorName = existingAuthor[0].AuthorName
	}
	if updateAuthor.AuthorAddress == "" {
		updateAuthor.AuthorAddress = existingAuthor[0].AuthorAddress
	}
	if updateAuthor.AuthorPhone == "" {
		updateAuthor.AuthorPhone = existingAuthor[0].AuthorPhone
	}

	if err := controller.AuthorSvc.UpdateAuthor(updateAuthor); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusCreated, "AuthorDetail is updated successfully")
}

// DeleteAuthor implements IAuthorController.
func (controller *AuthorController) DeleteAuthor(e echo.Context) error {
	tempAuthorID := e.Param("authorID")
	AuthorID, err := strconv.ParseInt(tempAuthorID, 0, 0)
	if err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid Data")
	}	
	_, err = controller.AuthorSvc.GetAuthors(uint(AuthorID))
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	if err := controller.AuthorSvc.DeleteAuthor(uint(AuthorID)); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	if err := controller.BookSvc.DeleteBookByAuthorID(uint(AuthorID)); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, "AuthorDetail is deleted successfully and All Books of the author deleted successfully")
}
