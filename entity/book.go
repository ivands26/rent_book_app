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
	//Rent     []Rent `gorm:"foreignKey:ID;"` //Buku bisa di pinjem banyak user
	Users []User `gorm:"many2many:user_books;"`
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

/*func (ab *AksesBook) GetDataBook() []Book {
	var daftarBook = []Book{}
	err := ab.DB.Model(&User{}).Select("Judul,ISBN,Author,Owned_by").Joins("left join emails on emails.user_id = users.id").Find(&daftarBook)
	if err.Error != nil {
	  log.Fatal(err.Statement.SQL.String())
	  return nil
	}

	return daftarBook
}
*/

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

func (ab *AksesBook) GetMyBook(emailUser string) []Book {
	var result = []Book{}
	err := ab.DB.Model(&Book{}).Select("books.ID, books.Judul, books.ISBN, books.Author, users.Nama").Joins("left join users on users.id = books.owned_by").Where("Email = ?", emailUser).Scan(&result)
	if err.Error != nil {
		log.Fatal(err.Statement.SQL.String())
		return nil
	}

	return result
}
