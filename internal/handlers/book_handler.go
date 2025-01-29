package handlers

import (
	"homework/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	service *services.BookService
}

func NewBookHandler(service *services.BookService) *BookHandler {
	return &BookHandler{service: service}
}

func (h *BookHandler) GetAllBooksHandler(c *gin.Context){
	books := h.service.GetAllBooks()
	c.JSON(http.StatusOK , books)
}

func (h *BookHandler) AddBookHandler(c *gin.Context){
	var input struct{
		Title string `json="title"`
		Author string `json="author"`
	}
	if err := c.ShouldBindBodyWithJSON(&input) ; err != nil{
		c.JSON(http.StatusBadRequest , gin.H{"error" : err.Error()})
	}
	book , err := h.service.AddBook(input.Title , input.Author)
	if err != nil{
		c.JSON(http.StatusBadRequest , gin.H{"error" : err.Error()})

	} 
	c.JSON(http.StatusOK , book)
}
