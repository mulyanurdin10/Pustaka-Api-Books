package handler

import (
	"fmt"
	"net/http"
	"pustaka-api/book"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

func (h *bookHandler) FindAll(ctx *gin.Context) {
	books, err := h.bookService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	responseBooks := []book.ResponseBook{}

	for _, v := range books {
		responseBook := convertToResponseBook(v)

		responseBooks = append(responseBooks, responseBook)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": responseBooks,
	})
}

func (h *bookHandler) FindById(ctx *gin.Context) {
	id := ctx.Param("id")
	idBook, _ := strconv.Atoi(id)
	book, err := h.bookService.FindById(idBook)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": convertToResponseBook(book),
	})
}

func (h *bookHandler) PostBooksHandler(ctx *gin.Context) {
	var requestBook book.RequestBook
	err := ctx.ShouldBindJSON(&requestBook)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}
	book, err := h.bookService.Create(requestBook)
	ctx.JSON(http.StatusCreated, gin.H{
		"data": convertToResponseBook(book),
	})
}

func (h *bookHandler) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	idBook, _ := strconv.Atoi(id)
	var updateBook book.UpdateBook
	err := ctx.ShouldBindJSON(&updateBook)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	book, err := h.bookService.Update(idBook, updateBook)
	ctx.JSON(http.StatusCreated, gin.H{
		"data": convertToResponseBook(book),
	})
}

func (h *bookHandler) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	idBook, _ := strconv.Atoi(id)
	book, err := h.bookService.Delete(idBook)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": convertToResponseBook(book),
	})
}

func convertToResponseBook(books book.Books) book.ResponseBook {
	return book.ResponseBook{
		ID:          books.ID,
		Title:       books.Title,
		Price:       books.Price,
		Discount:    books.Discount,
		Rating:      books.Rating,
		Description: books.Description,
	}
}
