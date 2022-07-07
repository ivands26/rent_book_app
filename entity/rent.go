package entity

import (
	"log"

	"gorm.io/gorm"
)

type Rent struct {
	gorm.Model
	IDUser uint
	IDBuku uint
}

type AksesRent struct {
	DB *gorm.DB
}

func (ar *AksesRent) RentBuku(userid, bookid uint) bool {
	daftarrent := Rent{IDUser: userid, IDBuku: bookid}
	err := ar.DB.Create(&daftarrent).Error
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func (ar *AksesRent) ReturnBuku(bookid uint) bool {
	deleteExc := ar.DB.Where("id_buku = ?", bookid).Delete(&Rent{})
	if err := deleteExc.Error; err != nil {
		log.Fatal(err)
		return false
	}
	if aff := deleteExc.RowsAffected; aff < 1 {
		log.Println("Gagal Mengembalikan Buku")
		return false
	}
	return true
}
