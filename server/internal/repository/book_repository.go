package repository

import (
	"github.com/nanonymoussu/bookstore/internal/models"
	"gorm.io/gorm"
)

type BookRepository interface {
	List(page, pageSize int) ([]models.Book, int64, error)
	GetByID(id int64) (models.Book, error)
	Create(book models.Book) (models.Book, error)
	Update(book models.Book) (models.Book, error)
	Delete(id int64) error
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db: db}
}

func (r *bookRepository) List(page, pageSize int) ([]models.Book, int64, error) {
	var books []models.Book
	var count int64

	offset := (page - 1) * pageSize

	if err := r.db.Model(&models.Book{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.Offset(offset).Limit(pageSize).Find(&books).Error; err != nil {
		return nil, 0, err
	}

	return books, count, nil
}

func (r *bookRepository) GetByID(id int64) (models.Book, error) {
	var book models.Book
	if err := r.db.First(&book, id).Error; err != nil {
		return book, err
	}

	return book, nil
}

func (r *bookRepository) Create(book models.Book) (models.Book, error) {
	if err := r.db.Create(&book).Error; err != nil {
		return book, err
	}

	return book, nil
}

func (r *bookRepository) Update(book models.Book) (models.Book, error) {
	if err := r.db.Save(&book).Error; err != nil {
		return book, err
	}

	return book, nil
}

func (r *bookRepository) Delete(id int64) error {
	return r.db.Delete(&models.Book{}, id).Error
}
