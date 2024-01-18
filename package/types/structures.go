package types

import (
	validate "github.com/go-ozzo/ozzo-validation"
)

// Response struct || marshalled into JSON format from struct
type BookRequest struct{
	BookID uint `json:"book_id,omitempty"`
	BookName string `json:"book_name"`
	AuthorId uint `json:"author_id"`
	Publication string `json:"publication,omitempty"`
}

func (book BookRequest) Validate() error{
	return validate.ValidateStruct(&book,
		validate.Field(&book.BookName, validate.Required.Error("Book name is required"),
		validate.Length(1,50)),
		validate.Field(&book.AuthorId, validate.Required.Error("Author id is required")),
	)
}


// Response struct || marshalled into JSON format from struct
type AuthorRequest struct{
	AuthorID uint `json:"author_id"`
	AuthorName string `json:"author_name"`
	AuthorAddress string `json:"author_address,omitempty"`
	AuthorPhone string `json:"author_phone,omitempty"`
}


func (author AuthorRequest) Validate() error{
	return validate.ValidateStruct(&author,
		validate.Field(&author.AuthorName, validate.Required.Error("Author name is required"),
		validate.Length(1,50)),
	)
}