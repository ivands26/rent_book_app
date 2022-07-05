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

func (au *AksesUser) GetUserEmail(EmailUser string) bool {
	getEmail := au.DB.Where("Email = ?", EmailUser).Find(&User{})
	if err := getEmail.Error; err != nil {
		return false
	}
	if aff := getEmail.RowsAffected; aff < 1 {
		return false
	}
	return true
}

func (au *AksesUser) GetUserPass(EmailUser, PassUser string) bool {
	getPass := au.DB.Where("Email = ? and Password = ?", EmailUser, PassUser).First(&User{})
	if err := getPass.Error; err != nil {
		return false
	}
	return true
}

func (as *AksesUser) GetAllData() []User {
	var daftarUser = []User{}
	err := as.DB.Find(&daftarUser)
	if err.Error != nil {
		log.Fatal(err.Statement.SQL.String())
		return nil
	}
	return daftarUser
}
