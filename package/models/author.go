package models


type AuthorDetails struct {
	AuthorId uint `gorm:"primaryKey;autoIncrement"`
	AuthorName string `json:"author_name"`
	AuthorAddress string `json:"author_address"`
	AuthorPhone string `json:"author_phone"`
}