package entity

import (
	"log"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nama     string
	No_hp    string
	Email    string
	Password string
}

type AksesUser struct {
	DB *gorm.DB
}

func (au *AksesUser) RegisUser(newUser User) User {
	err := au.DB.Create(&newUser).Error
	if err != nil {
		log.Fatal(err)
		return User{}
	}

	return newUser
}

func (au *AksesUser) GetDataUser() []User {
	var daftarUser = []User{}
	err := au.DB.Find(&daftarUser)
	if err.Error != nil {
		log.Fatal(err.Statement.SQL.String())
		return nil
	}

	return daftarUser
}
