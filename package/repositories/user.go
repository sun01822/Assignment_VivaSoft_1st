package repositories

import (
	"Assignment_Vivasoft/package/domain"
	"Assignment_Vivasoft/package/models"
	"gorm.io/gorm"
)

// parent struct to implement interface binding
type userRepo struct {
	db *gorm.DB
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
	err := repo.db.Create(user).Error
	if err != nil {
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



