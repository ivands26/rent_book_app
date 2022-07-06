package entity

import (
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ISBN     string
	Judul    string
	Author   string
	Owned_by uint
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
	uid := uuid.New()
	newBook.ISBN = uid.String()
	err := ab.DB.Create(&newBook).Error
	if err != nil {
		log.Fatal(err)
		return Book{}
	}

	return newBook
}

func (ab *AksesBook) DeleteBook(IDBook int) bool {
	postExc := ab.DB.Where("ID = ?", IDBook).Delete(&Book{})
	if err := postExc.Error; err != nil {
		log.Fatal(err)
		return false
	}
	if aff := postExc.RowsAffected; aff < 1 {
		log.Println("Gagal menghapus buku")
	}
	return true
}
