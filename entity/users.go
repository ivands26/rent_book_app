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

func (au *AksesUser) GetUserPass(PassUser string) bool {
	getPass := au.DB.Where("Password = ?", PassUser).Find(&User{})
	if err := getPass.Error; err != nil {
		return false
	}
	if aff := getPass.RowsAffected; aff < 1 {
		return false
	}
	return true
}

func (au *AksesUser) GetName(EmailUser string) []User {
	var daftarUser = []User{}
	err := au.DB.Where("Email = ?", EmailUser).Select("Nama").Find(&daftarUser)
	if err.Error != nil {
		log.Fatal(err.Statement.SQL.String())
		return nil
	}
	return daftarUser
}

func (au *AksesUser) GetEmailPass(EmailUser, PassUser string) bool {
	getPass := au.DB.Where("Email = ? and Password = ?", EmailUser, PassUser).First(&User{})
	if err := getPass.Error; err != nil {
		return false
	}
	return true
}

func (au *AksesUser) GetProfileUser(EmailUser string) []User {
	var profileUser = []User{}
	err := au.DB.Where("Email = ?", EmailUser).Find(&profileUser)
	if err.Error != nil {
		log.Fatal(err.Statement.SQL.String())
		return nil
	}

	return profileUser
}

func (au *AksesUser) DeleteUser(EmailUser string) bool {
	postExc := au.DB.Where("Email = ?", EmailUser).Delete(&User{})

	if err := postExc.Error; err != nil {
		log.Fatal(err)
		return false
	}

	if aff := postExc.RowsAffected; aff < 1 {
		log.Println("Gagal Menghapus Akun")
		return false
	}

	return true
}

func (au *AksesUser) GetAllData() []User {
	var daftarUser = []User{}
	err := au.DB.Find(&daftarUser)
	if err.Error != nil {
		log.Fatal(err.Statement.SQL.String())
		return nil
	}
	return daftarUser
}
