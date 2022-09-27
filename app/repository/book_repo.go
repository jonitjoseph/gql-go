package repository

import (
	"gql-go/app/models"
	"gql-go/graph/model"

	"gorm.io/gorm"
)

type BookRepository interface {
	CreateBook(bookInput *model.BookInput) (*models.Book, error)
	UpdateBook(bookInput *model.BookInput, id int) error
	DeleteBook(id int) error
	GetOneBook(id int) (*models.Book, error)
	GetAllBooks() ([]*model.Book, error)
}

type BookService struct {
	Db *gorm.DB
}

// CreateBook implements BookRepository
func (b *BookService) CreateBook(bookInput *model.BookInput) (*models.Book, error) {
	book := &models.Book{
		Title:     bookInput.Title,
		Author:    bookInput.Author,
		Publisher: bookInput.Publisher,
	}
	err := b.Db.Create(&book).Error

	return book, err
}

// DeleteBook implements BookRepository
func (b *BookService) DeleteBook(id int) error {
	book := &models.Book{}
	err := b.Db.Delete(book, id).Error
	return err
}

// GetAllBooks implements BookRepository
func (b *BookService) GetAllBooks() ([]*model.Book, error) {
	books := []*model.Book{}
	err := b.Db.Find(&books).Error
	return books, err
}

// GetOneBook implements BookRepository
func (b *BookService) GetOneBook(id int) (*models.Book, error) {
	book := &models.Book{}
	err := b.Db.Where("id = ?", id).First(book).Error
	return book, err
}

// UpdateBook implements BookRepository
func (b *BookService) UpdateBook(bookInput *model.BookInput, id int) error {
	book := models.Book{
		ID:        id,
		Title:     bookInput.Title,
		Author:    bookInput.Author,
		Publisher: bookInput.Publisher,
	}
	err := b.Db.Model(&book).Where("id = ?", id).Updates(book).Error
	return err
}

var _ BookRepository = &BookService{}

func NewBookService(db *gorm.DB) *BookService {
	return &BookService{
		Db: db,
	}
}
