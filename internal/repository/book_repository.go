package repository

import "homework/internal/models"

type BookRepository struct {
	books  []models.Book
	nextID int
}

func NewRepository() *BookRepository {
	return &BookRepository{
		books:  []models.Book{},
		nextID: 1,
	}
}

func (r *BookRepository) GetAllBooks() []models.Book {
	return r.books
}

func (r *BookRepository) AddBook(book models.Book) models.Book {
	book.ID = r.nextID
	r.nextID++
	r.books = append(r.books, book)
	return book
}
