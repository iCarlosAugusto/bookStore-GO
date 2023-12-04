package main

import (
	"github.com/gin-gonic/gin"
	"github.com/iCarlosAugusto/Golang/books/v2/controller"
	"github.com/iCarlosAugusto/Golang/books/v2/repositories"
	routers "github.com/iCarlosAugusto/Golang/books/v2/router"
	"github.com/iCarlosAugusto/Golang/common/database"
)

func main() {

	router := gin.Default();
	dbHandler := database.Init("postgres://ejelfjhi:nG_smKv6uovq1GdmNZu9eRg6p0NU8nQT@berry.db.elephantsql.com/ejelfjhi");

	bookRepositoryTeste := repositories.BookRepositoryImp{Db: dbHandler}

	bookController := controller.BookController{ BookRepository: bookRepositoryTeste}
	routers.InitRouters(&bookController, router);

	router.Run("localhost:3030")
}
