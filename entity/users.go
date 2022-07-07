package entity

import (
	"log"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nama     string
	No_hp    string
	Email    string `gorm:"unique"`
	Password string
	Books    []Book `gorm:"foreignKey:Owned_by;OnUpdate:CASCADE,OnDelete:SET NULL"` //user punya banyak buku
	// Rent     []Rent `gorm:"foreignKey:User_id"`
	//Rent     []Rent `gorm:"foreignKey:ID;"` //user bisa pinjem banyak buku
	//Buku []Book `gorm:"many2many:user_books;"`
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

func (au *AksesUser) GetEmailPass(EmailUser, PassUser string) bool {
	getPass := au.DB.Where("Email = ? and Password = ?", EmailUser, PassUser).First(&User{})
	if err := getPass.Error; err != nil {
		return false
	}
	return true
}

func (au *AksesUser) GetProfileUser(EmailUser string) User {
	var profileUser = User{}
	err := au.DB.Where("Email = ?", EmailUser).First(&profileUser)
	if err.Error != nil {
		log.Fatal(err.Statement.SQL.String())
		return profileUser
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

func (au *AksesUser) UpdateUserNama(emailUser string, namaUpdate string) bool {
	updateExc := au.DB.Model(&User{}).Where("Email = ?", emailUser).Update("Nama", namaUpdate)
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

func (au *AksesUser) UpdateUserNo(emailUser string, hpUpdate string) bool {
	updateExc := au.DB.Model(&User{}).Where("Email = ?", emailUser).Update("No_hp", hpUpdate)
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

func (au *AksesUser) UpdateUserSurel(emailUser string, surelUpdate string) bool {
	updateExc := au.DB.Model(&User{}).Where("Email = ?", emailUser).Update("Email", surelUpdate)
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

func (au *AksesUser) UpdateUserPass(emailUser string, PassUpdate string) bool {
	updateExc := au.DB.Model(&User{}).Where("Email = ?", emailUser).Update("Password", PassUpdate)
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

// func UpdateRent(au *AksesUserBook) (emailUser string, idbuku int) {
// 	var daftarBook []Book
// 	updateExc := au.DB.Preload("User").Where("Email = ?", emailUserFind(&daftarBook)
// 	if err := updateExc.Error; err != nil {
// 		log.Fatal(err)
// 		return false
// 	}
// 	if aff := updateExc.RowsAffected; aff < 1 {
// 		log.Println("Tidak ada data yang diupdate")
// 		return false
// 	}
// 	return true
// }
