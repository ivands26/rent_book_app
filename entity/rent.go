package entity

// import (
// 	"log"

// 	"gorm.io/gorm"
// )

// type Rent struct {
// 	gorm.Model
// 	User_id uint
// 	Book_id uint
// }

// type AksesRent struct {
// 	DB *gorm.DB
// }

// func (ar *AksesRent) PinjemBuku(newRent Rent) Rent {
// 	err := ar.DB.Create(&newRent).Error
// 	if err != nil {
// 		log.Fatal(err)
// 		return Rent{}
// 	}
// 	return newRent

// }
