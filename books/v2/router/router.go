package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iCarlosAugusto/Golang/books/v2/controller"
)

func InitRouters(bookController *controller.BookController, ginContext *gin.Engine) {

	ginContext.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, "welcome home")
	});

	ginContext.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	router := ginContext.Group("/books")
	
	router.POST("", bookController.Save);
	router.GET("", bookController.FindAll);
}