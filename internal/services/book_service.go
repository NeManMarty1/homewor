package services

import (
	"errors"
	"homework/internal/models"
	"homework/internal/repository"
)

type BookService struct {
	repo *repository.BookRepository
}

func NewBookService(repo *repository.BookRepository) *BookService {
	return &BookService{repo: repo}
}

func (b *BookService) GetAllBooks() []models.Book {
	return b.repo.GetAllBooks()
}

func (b *BookService) AddBook(title, author string) (models.Book, error) {
	if title == "" || author == "" {
		return models.Book{}, errors.New("title or author are requared")
	}
	newBook := models.Book{
		Title:  title,
		Author: author,
	}
	return b.repo.AddBook(newBook), nil
}
