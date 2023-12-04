package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iCarlosAugusto/Golang/books/v2/helpers"
	"github.com/iCarlosAugusto/Golang/books/v2/repositories"
	"github.com/iCarlosAugusto/Golang/common/models"
)

type SaveBookRequestBody struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

type BookController struct {
	BookRepository repositories.BookRepositoryImp
}


func (b BookController) Save(ctx *gin.Context)  {

	createBookRequest := SaveBookRequestBody{};

	if err := ctx.BindJSON(&createBookRequest); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	book := models.Book{
		Title: createBookRequest.Title,
		Author: createBookRequest.Author,
		Description: createBookRequest.Description,
	}

	b.BookRepository.Save(book);

	ctx.JSON(http.StatusCreated, &book);
}

func (b BookController) FindAll(ctx *gin.Context) {
	result := b.BookRepository.FindAll();
	webResponse := helpers.WebResponse{
		Code: 200,
		Status: "OK",
		Data: result,
	}
	ctx.JSON(http.StatusOK, webResponse);
	
}