package controllers

import (
	"Assignment_Vivasoft/package/domain"
	"Assignment_Vivasoft/package/models"
	"Assignment_Vivasoft/package/types"
	"net/http"
	"strconv"
	"time"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type IUserController interface {
	GetUsers(e echo.Context) error
	CreateUser(e echo.Context) error
	UpdateUser(e echo.Context) error
	DeleteUser(e echo.Context) error
}

type UserController struct {
	UserSvc domain.IUserService
}

func NewUserController(UserSvc domain.IUserService) UserController {
	return UserController{
		UserSvc: UserSvc,
	}
}


// CreateUser implements IUserController.
func (controller *UserController) CreateUser(e echo.Context) error {
	reqUser := &types.UserRequest{}
	if err := e.Bind(reqUser); err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid Data")
	}
	if err := reqUser.Validate(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	User := &models.UserDetails{
		UserName:    reqUser.UserName,
		Email:   reqUser.Email,
		Password:    reqUser.Password,
		Phone: reqUser.Phone,
		Address: reqUser.Address,
	}
	if err := controller.UserSvc.CreateUser(User); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusCreated, "UserDetail is created successfully")
}

// GetUsers implements IUserController.
func (controller *UserController) GetUsers(e echo.Context) error {
	tempUserID := e.QueryParam("userID")
	UserID, err := strconv.ParseInt(tempUserID, 0, 0)
	if err != nil && tempUserID != "" {
		return e.JSON(http.StatusBadRequest, "Enter a valid User ID")
	}
	User, err := controller.UserSvc.GetUsers(&gorm.Model{ID: uint(UserID)})
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, User)
}

// UpdateUser implements IUserController.
func (controller *UserController) UpdateUser(e echo.Context) error {
	//reqUser := &types.UserRequest{}
	reqUser := &types.UserUpdateRequest{}
	if err := e.Bind(reqUser); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	// if err := reqUser.Validate(); err != nil {
	// 	return e.JSON(http.StatusBadRequest, err.Error())
	// }
	tempUserID := e.Param("userID")
	UserID, err := strconv.ParseInt(tempUserID, 0, 0)
	if err != nil {
		return e.JSON(http.StatusBadRequest, "Enter a valid User ID")
	}
	existingUser, err := controller.UserSvc.GetUsers(&gorm.Model{ID: uint(UserID)})
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	
	updateUser := &models.UserDetails{
		Model: gorm.Model{ID: uint(UserID), UpdatedAt: time.Now(), CreatedAt: time.Now()},
		UserName:  		reqUser.UserName,
		Email: 			reqUser.Email,
		Password:   	reqUser.Password,
		Phone: 			reqUser.Phone,
		Address: 		reqUser.Address,
	}

	if updateUser.UserName == "" {
		updateUser.UserName = existingUser[0].UserName
	}
	if updateUser.Email == "" {
		updateUser.Email = existingUser[0].Email
	}
	if updateUser.Password == "" {
		updateUser.Password = existingUser[0].Password
	}
	if updateUser.Phone == "" {
		updateUser.Phone = existingUser[0].Phone
	}
	if updateUser.Address == "" {
		updateUser.Address = existingUser[0].Address
	}

	if err := controller.UserSvc.UpdateUser(updateUser); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusCreated, "UserDetail is updated successfully")
}

// DeleteUser implements IUserController.
func (controller *UserController) DeleteUser(e echo.Context) error {
	tempUserID := e.Param("userID")
	UserID, err := strconv.ParseInt(tempUserID, 0, 0)
	if err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid Data")
	}	
	_, err = controller.UserSvc.GetUsers(&gorm.Model{ID: uint(UserID)})
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	if err := controller.UserSvc.DeleteUser(&gorm.Model{ID: uint(UserID)}); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, "UserDetail is deleted successfully")
}
