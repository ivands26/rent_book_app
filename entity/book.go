package entity

import (
	"log"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Judul  string
	Author string
}

type AksesBook struct {
	DB *gorm.DB
}

func (ab *AksesBook) GetDataBook() []Book {
	var daftarBook = []Book{}
	err := ab.DB.Find(&daftarBook)
	if err.Error != nil {
		log.Fatal(err.Statement.SQL.String())
		return nil
	}

	return daftarBook
}

func (ab *AksesBook) InputBook(newBook Book) Book {
	err := ab.DB.Create(&newBook).Error
	if err != nil {
		log.Fatal(err)
		return Book{}
	}

	return newBook
}
