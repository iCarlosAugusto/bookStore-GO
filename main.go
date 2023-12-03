package main

import (
	"github.com/gin-gonic/gin"
	"github.com/iCarlosAugusto/Golang/books"
	"github.com/iCarlosAugusto/Golang/common/database"
)

func main() {

	router := gin.Default();
	dbHandler := database.Init("postgres://ejelfjhi:nG_smKv6uovq1GdmNZu9eRg6p0NU8nQT@berry.db.elephantsql.com/ejelfjhi");
	books.RegisterRoutes(router, dbHandler);

	
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
				"port":  "3000",
				"dbUrl": "postgres://ejelfjhi:nG_smKv6uovq1GdmNZu9eRg6p0NU8nQT@berry.db.elephantsql.com/ejelfjhi",
		})	
	})

	router.Run("localhost:3000")
}
