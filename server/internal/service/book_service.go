package service

import (
	"github.com/nanonymoussu/bookstore/internal/models"
	"github.com/nanonymoussu/bookstore/internal/repository"
)

type BookService interface {
	ListBooks(page, pageSize int) ([]models.Book, int64, error)
	GetBook(id int64) (models.Book, error)
	CreateBook(title, author, isbn string, price float64, stock int32) (models.Book, error)
	UpdateBook(id int64, title, author, isbn string, price float64, stock int32) (models.Book, error)
	DeleteBook(id int64) error
}

type bookService struct {
	repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) BookService {
	return &bookService{repo: repo}
}

func (s *bookService) ListBooks(page, pageSize int) ([]models.Book, int64, error) {
	return s.repo.List(page, pageSize)
}

func (s *bookService) GetBook(id int64) (models.Book, error) {
	return s.repo.GetByID(id)
}

func (s *bookService) CreateBook(title, author, isbn string, price float64, stock int32) (models.Book, error) {
	book := models.Book{
		Title:  title,
		Author: author,
		ISBN:   isbn,
		Price:  price,
		Stock:  stock,
	}

	return s.repo.Create(book)
}

func (s *bookService) UpdateBook(id int64, title, author, isbn string, price float64, stock int32) (models.Book, error) {
	book, err := s.repo.GetByID(id)
	if err != nil {
		return book, err
	}

	book.Title = title
	book.Author = author
	book.ISBN = isbn
	book.Price = price
	book.Stock = stock

	return s.repo.Update(book)
}

func (s *bookService) DeleteBook(id int64) error {
	return s.repo.Delete(id)
}
