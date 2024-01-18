package services

import (
	"Assignment_Vivasoft/package/domain"
	"Assignment_Vivasoft/package/models"
	"Assignment_Vivasoft/package/types"
	"errors"
)

// Parent struct to implement interface binding
type bookService struct {
	repo domain.IBookRepo
}

// Interface binding
func BookServiceInstance(bookRepo domain.IBookRepo) domain.IBookService {
	return &bookService{
		repo: bookRepo,
	}
}

// All methods of interface are implemented here
// GetBooks implements domain.IBookService.
func (service *bookService) GetBooks(bookID uint) ([]types.BookRequest, error) {
	var allBooks []types.BookRequest

	book := service.repo.GetBooks(bookID)
	if len(book) == 0 {
		return nil, errors.New("no books found")
	}
	for _, val := range book {
		allBooks = append(allBooks, types.BookRequest{
			BookID:       val.BookId,
			BookName:     val.BookName,
			AuthorId:     val.AuthorId,
			Publication:  val.Publication,
		})
	}
	return allBooks, nil
}

// CreateBook implements domain.IBookService.
func (service *bookService) CreateBook(book *models.BookDetails) error {
	if err := service.repo.CreateBook(book); err != nil {
		return errors.New("bookdetail is not created")
	}
	return nil
}

// UpdateBook implements domain.IBookService.
func (service *bookService) UpdateBook(book *models.BookDetails) error {
	if err := service.repo.UpdateBook(book); err != nil {
		return errors.New("bookdetail is not updated")
	}
	return nil
}

// DeleteBook implements domain.IBookService.
func (service *bookService) DeleteBook(bookID uint) error {
	if err := service.repo.DeleteBook(bookID); err != nil {
		return errors.New("bookdetail is not deleted")
	}
	return nil
}

// DeleteBookByAuthorID implements domain.IBookService
func (service *bookService) DeleteBookByAuthorID(authorID uint) error {
	if err := service.repo.DeleteBookByAuthorID(authorID); err != nil {
		return errors.New("bookdetail is not deleted")
	}
	return nil
}

