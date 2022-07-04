package main

import (
	"fmt"

	"group_project/config"
	"group_project/entity"
)

func main() {
	conn := config.InitDB()
	config.MigrateDB(conn)
	aksesUser := entity.AksesUser{DB: conn}

	var input int
	for input != 99 {
		fmt.Println("\tRENT BOOKS APP")
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Println("3. Lihat Koleksi Buku")
		fmt.Println("99. Keluar")
		fmt.Println("Masukkan Input Anda :")
		fmt.Scanln(&input)

		switch input {
		case 1:
			var newUser entity.User
			fmt.Println("Masukan Nama")
			fmt.Scanln(&newUser.Nama)
			fmt.Println("Masukan No HP")
			fmt.Scanln(&newUser.No_hp)
			fmt.Println("Masukan Email")
			fmt.Scanln(&newUser.Email)
			fmt.Println("Masukan Password")
			fmt.Scanln(&newUser.Password)
			res := aksesUser.RegisUser(newUser)
			if res.ID == 0 {
				fmt.Println("Registrasi Gagal")
				break
			}
			fmt.Println("Registrasi Berhasil")

		case 2:
			fmt.Println("Daftar User")
			for _, value := range aksesUser.GetDataUser() {
				fmt.Println(value.Nama)
				fmt.Println(value.No_hp)
				fmt.Println(value.Email)
				fmt.Println(" ")
			}

		}
	}
	fmt.Println("TERIMA KASIH TELANG MENGGUNAKAN APLIKASI KAMI")
}
