package repositories

import (
	"Assignment_Vivasoft/package/domain"
	"Assignment_Vivasoft/package/models"
	"errors"
	"gorm.io/gorm"
	"Assignment_Vivasoft/package/utils"
)

// parent struct to implement interface binding
type userRepo struct {
	db *gorm.DB
}

// LoginUser implements domain.IUserRepo.
func (repo *userRepo) LoginUser(user *models.UserDetails) error {
	// Find the user by user_name
	var existingUser models.UserDetails
	if err := repo.db.Where("user_name = ?", user.UserName).First(&existingUser).Error; err != nil {
		return err
	}
	
	// Compare the stored hashed password, with the hashed version of the password that was received
	if err := utils.ComparePassword(existingUser.Password, user.Password); err != nil {
		return err
	}
	// Passwords match, authentication successful
	return nil
}


// GetUsers implements domain.IUserRepo.
func (repo *userRepo) GetUsers(model *gorm.Model) []models.UserDetails {
	var user []models.UserDetails
	var err error
	userID := model.ID

	if userID != 0 {
		err = repo.db.Where("id = ?", userID).Find(&user).Error
	} else {
		err = repo.db.Find(&user).Error
	}
	if err != nil {
		return []models.UserDetails{}
	}
	return user
}

// CreateUser implements domain.IUserRepo.
func (repo *userRepo) CreateUser(user *models.UserDetails) error {
	userName := user.UserName

	// Check if the user with the same userName already exists
	var existingUser models.UserDetails
	err := repo.db.Where("user_name = ?", userName).First(&existingUser).Error
	if err == nil {
		// User with the same userName already exists, return an error
		return errors.New("user with the same userName already exists")
	}

	// hashedPassword
	user.Password = utils.HashPassword(user.Password)

	err2 := repo.db.Create(user).Error
	if err2 != nil {
		return err
	}
	return nil
}


// DeleteUser implements domain.IUserRepo.
func (repo *userRepo) DeleteUser(model *gorm.Model) error {
	var user models.UserDetails
	if err := repo.db.Where("id = ?", model.ID).Delete(&user).Error; err != nil {
		return err
	}
	return nil
}

// UpdateUser implements domain.IUserRepo.
func (repo *userRepo) UpdateUser(user *models.UserDetails) error {
	err := repo.db.Save(user).Error
	if err != nil {
		return err
	}
	return nil
}

// interface binding
func UserDBInstance(d *gorm.DB) domain.IUserRepo {
	return &userRepo{
		db: d,
	}
}
