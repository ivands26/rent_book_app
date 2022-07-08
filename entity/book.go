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
	Status   string
	Rent     []Rent `gorm:"foreignKey:IDBuku;"`
}

type Joinbuku struct {
	ID       uint
	ISBN     string
	Judul    string
	Author   string
	Nama     string
	Owned_by int
	Status   string
}

type AksesBook struct {
	DB *gorm.DB
}

func (ab *AksesBook) GetDataBook() []Joinbuku {
	var result = []Joinbuku{}
	err := ab.DB.Model(&Book{}).Select("books.ID, books.Judul, books.ISBN, books.Author, users.Nama, books.Status").Joins("left join users on users.id = books.owned_by").Scan(&result)
	if err.Error != nil {
		log.Fatal(err.Statement.SQL.String())
		return nil
	}
	return result
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

func (ab *AksesBook) GetMyBook(emailUser string) []Joinbuku {
	var result = []Joinbuku{}
	err := ab.DB.Model(&Book{}).Select("books.ID, books.Judul, books.ISBN, books.Author, users.Nama, books.Status").Joins("left join users on users.id = books.owned_by").Where("Email = ?", emailUser).Scan(&result)
	if err.Error != nil {
		log.Fatal(err.Statement.SQL.String())
		return nil
	}
	return result
}

func (ab *AksesBook) UpdateBook(emailUser string, judulUpdate string, authorUpdate string) bool {
	updateExc := ab.DB.Model(&Book{}).Where("Email = ?", emailUser).Updates(Book{Judul: judulUpdate, Author: authorUpdate})
	if err := updateExc.Error; err != nil {
		log.Fatal(err)
		return false
	}
	if aff := updateExc.RowsAffected; aff < 1 {
		log.Println("Tidak ada data yang diupdate")
		return false
	}
	return true
}

func (ab *AksesBook) GetBookJA(emailUser string) []Book {
	var daftarBook = []Book{}
	err := ab.DB.Where("Email = ?", emailUser).Find(&daftarBook)
	if err.Error != nil {
		log.Fatal(err.Statement.SQL.String())
		return nil
	}
	return daftarBook
}

func (ab *AksesBook) UpdateBookJA(id int, judulUpdate string, authorUpdate string) bool {
	kosong := map[string]interface{}{}
	if judulUpdate != "" {
		kosong["judul"] = judulUpdate
	}
	if authorUpdate != "" {
		kosong["author"] = authorUpdate
	}
	updateExc := ab.DB.Model(&Book{}).Where("ID = ?", id).Updates(Book{Judul: judulUpdate, Author: authorUpdate})
	if err := updateExc.Error; err != nil {
		log.Fatal(err)
		return false
	}
	if aff := updateExc.RowsAffected; aff < 1 {
		log.Println("Tidak ada data yang diupdate")
		return false
	}
	return true
}

func (ab *AksesBook) GetDataRentBook(emailUser string, bookstat string) []Joinbuku { //kurang yg udah dipinjam org
	var daftarBook = []Joinbuku{}
	err := ab.DB.Model(&Book{}).Select("books.ID, books.Judul, books.ISBN, books.Author, users.Nama").Joins("left join users on users.id = books.owned_by").Where("Email != ? and books.Status = ?", emailUser, bookstat).Scan(&daftarBook)
	if err.Error != nil {
		log.Fatal(err.Statement.SQL.String())
		return nil
	}

	return daftarBook
}

func (ab *AksesBook) GetDataYourRentedBook(userid uint) []Joinbuku {
	var daftarrent = []Joinbuku{}
	err := ab.DB.Model(&Rent{}).Select("books.ID, books.Judul, books.ISBN, books.Author, books.Owned_by").Joins("left join books on books.ID = rents.id_buku").Joins("left join users on users.id = rents.id_user").Where("users.id = ?", userid).Find(&daftarrent)
	if err.Error != nil {
		log.Fatal(err.Statement.SQL.String())
		return nil
	}
	return daftarrent
}

func (ab *AksesBook) Updatestatus(bookid uint, stat string) bool {
	updateExc := ab.DB.Model(&Book{}).Where("books.ID = ?", bookid).Update("books.Status", stat)
	if err := updateExc.Error; err != nil {
		log.Fatal(err)
		return false
	}
	if aff := updateExc.RowsAffected; aff < 1 {
		log.Println("Tidak ada data yang diupdate")
		return false
	}
	return true
}
