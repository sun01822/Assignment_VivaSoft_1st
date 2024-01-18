package services

import (
	"Assignment_Vivasoft/package/domain"
	"Assignment_Vivasoft/package/models"
	"Assignment_Vivasoft/package/types"
	"errors"
)

// Parent struct to implement interface binding
type authorService struct {
	repo domain.IAuthorRepo
}

// Interface binding
func AuthorServiceInstance(authorRepo domain.IAuthorRepo) domain.IAuthorService {
	return &authorService{
		repo: authorRepo,
	}
}

// All methods of interface are implemented here
// GetBooks implements domain.IBookService.
func (service *authorService) GetAuthors(authorID uint) ([]types.AuthorRequest, error) {
	var allAuthors []types.AuthorRequest

	author := service.repo.GetAuthors(authorID)
	if len(author) == 0 {
		return nil, errors.New("no authors found")
	}
	for _, val := range author {
		allAuthors = append(allAuthors, types.AuthorRequest{
			AuthorID:          val.AuthorId,
			AuthorName:        val.AuthorName,
			AuthorAddress:     val.AuthorAddress,
			AuthorPhone:  	   val.AuthorPhone,
		})
	}
	return allAuthors, nil
}

// CreateBook implements domain.IBookService.
func (service *authorService) CreateAuthor(author *models.AuthorDetails) error {
	if err := service.repo.CreateAuthor(author); err != nil {
		return errors.New("authorDetail is not created")
	}
	return nil
}

// UpdateBook implements domain.IBookService.
func (service *authorService) UpdateAuthor(author *models.AuthorDetails) error {
	if err := service.repo.UpdateAuthor(author); err != nil {
		return errors.New("authordetail is not updated")
	}
	return nil
}

// DeleteBook implements domain.IBookService.
func (service *authorService) DeleteAuthor(authorID uint) error {
	if err := service.repo.DeleteAuthor(authorID); err != nil {
		return errors.New("authordetail is not deleted")
	}
	return nil
}

