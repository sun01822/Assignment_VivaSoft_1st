package repositories

import (
	"Assignment_Vivasoft/package/domain"
	"Assignment_Vivasoft/package/models"
	"gorm.io/gorm"
)

// parent struct to implement interface binding
type authorRepo struct{
	db *gorm.DB
}

// interface binding
func AuthorDBInstance(d *gorm.DB) domain.IAuthorRepo{
	return &authorRepo{
		db: d,
	}
}

// all methods of interface are implemented here
func (repo *authorRepo) GetAuthors(authorID uint)[]models.AuthorDetails{
	var author []models.AuthorDetails
	var err error

	if authorID != 0{
		err = repo.db.Where("author_id = ?", authorID).Find(&author).Error
	}else{
		err = repo.db.Find(&author).Error
	}
	if err != nil{
		return []models.AuthorDetails{}
	}
	return author
}


func (repo *authorRepo) CreateAuthor(author *models.AuthorDetails) error{
	err := repo.db.Create(author).Error
	if err != nil{
		return err
	}
	return nil
}

func (repo *authorRepo) UpdateAuthor(author *models.AuthorDetails) error{
	err := repo.db.Save(author).Error
	if err != nil{
		return err
	}
	return nil
}

func (repo *authorRepo) DeleteAuthor(authorID uint) error{
	var Author models.AuthorDetails
	if err := repo.db.Where("author_id = ?", authorID).Delete(&Author).Error; err != nil {
		return err
	}
	return nil
}
