package repositories

import (
	"github.com/iCarlosAugusto/Golang/common/models"
	"gorm.io/gorm"
)

type IBookReposotory interface {
	Save(book models.Book)
	FindAll() []models.Book
}

type BookRepositoryImp struct {
	Db *gorm.DB;
}

func (b BookRepositoryImp) Save(book models.Book) {
	b.Db.Create(&book);
}

func (b BookRepositoryImp) FindAll() []models.Book {
	var books []models.Book;
	b.Db.Find(&books);
	return books;
}