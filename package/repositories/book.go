package repositories

import (
	"Assignment_Vivasoft/package/domain"
	"Assignment_Vivasoft/package/models"
	"gorm.io/gorm"
)

// parent struct to implement interface binding
type bookRepo struct{
	db *gorm.DB
}

// interface binding
func BookDBInstance(d *gorm.DB) domain.IBookRepo{
	return &bookRepo{
		db: d,
	}
}

// all methods of interface are implemented here
func (repo *bookRepo) GetBooks(bookID uint)[]models.BookDetails{
	var book []models.BookDetails
	var err error

	if bookID != 0{
		err = repo.db.Where("book_id = ?", bookID).Find(&book).Error
	}else{
		err = repo.db.Find(&book).Error
	}
	if err != nil{
		return []models.BookDetails{}
	}
	return book
}


func (repo *bookRepo) CreateBook(book *models.BookDetails) error{
	err := repo.db.Create(book).Error
	if err != nil{
		return err
	}
	return nil
}

func (repo *bookRepo) UpdateBook(book *models.BookDetails) error{
	err := repo.db.Save(book).Error
	if err != nil{
		return err
	}
	return nil
}

func (repo *bookRepo) DeleteBook(bookID uint) error{
	var Book models.BookDetails
	if err := repo.db.Where("book_id = ?", bookID).Delete(&Book).Error; err != nil {
		return err
	}
	return nil
}

func (repo *bookRepo) DeleteBookByAuthorID(authorID uint) error{
	var Book models.BookDetails
	if err := repo.db.Where("author_id = ?", authorID).Delete(&Book).Error; err != nil {
		return err
	}
	return nil
}
