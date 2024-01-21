package services

import (
	"Assignment_Vivasoft/package/domain"
	"Assignment_Vivasoft/package/models"
	"Assignment_Vivasoft/package/types"
	"errors"
	"gorm.io/gorm"
)

// Parent struct to implement interface binding
type userService struct {
	repo domain.IUserRepo
}

// Interface binding
func UserServiceInstance(userRepo domain.IUserRepo) domain.IUserService {
	return &userService{
		repo: userRepo,
	}
}

// All methods of interface are implemented here
// GetBooks implements domain.IBookService.
func (service *userService) GetUsers(model *gorm.Model) ([]types.UserRequest, error) {
	var allusers []types.UserRequest
	user := service.repo.GetUsers(model)
	if len(user) == 0 {
		return nil, errors.New("no users found")
	}
	for _, val := range user {
		allusers = append(allusers, types.UserRequest{
			UserName:    	val.UserName,
			Email: 	 		val.Email,
			Password:    	val.Password,
			Phone:       	val.Phone,
			Address:    	val.Address,
		})
	}
	return allusers, nil
}

// CreateBook implements domain.IBookService.
func (service *userService) CreateUser(user *models.UserDetails) error {
	if err := service.repo.CreateUser(user); err != nil {
		return errors.New("userDetail is not created")
	}
	return nil
}

// UpdateBook implements domain.IBookService.
func (service *userService) UpdateUser(user *models.UserDetails) error {
	if err := service.repo.UpdateUser(user); err != nil {
		return errors.New("userdetail is not updated")
	}
	return nil
}

// DeleteBook implements domain.IBookService.
func (service *userService) DeleteUser(model *gorm.Model) error {
	if err := service.repo.DeleteUser(model); err != nil {
		return errors.New("userdetail is not deleted")
	}
	return nil
}

