package domain

import (
	"Assignment_Vivasoft/package/models"
	"Assignment_Vivasoft/package/types"
)

// for database Repository operation (call from service)
type IBookRepo interface{
	GetBooks(bookID uint)[]models.BookDetails
	CreateBook(book *models.BookDetails) error
	UpdateBook(book *models.BookDetails) error
	DeleteBook(bookID uint) error
	DeleteBookByAuthorID(authorID uint) error
}


// for service operation (response to controller || call from controller)
type IBookService interface{
	GetBooks(bookID uint) ([]types.BookRequest, error)
	CreateBook(book *models.BookDetails) error
	UpdateBook(book *models.BookDetails) error
	DeleteBook(bookID uint) error
	DeleteBookByAuthorID(authorID uint) error
}

// for database Repository operation (call from service)
type IAuthorRepo interface{
	GetAuthors(authorID uint)[]models.AuthorDetails
	CreateAuthor(author *models.AuthorDetails) error
	UpdateAuthor(author *models.AuthorDetails) error
	DeleteAuthor(authorID uint) error
}

// for service operation (response to controller || call from controller)
type IAuthorService interface{
	GetAuthors(authorID uint) ([]types.AuthorRequest, error)
	CreateAuthor(author *models.AuthorDetails) error
	UpdateAuthor(author *models.AuthorDetails) error
	DeleteAuthor(authorID uint) error
}