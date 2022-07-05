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
		fmt.Println("\n\tWelcome in RENT BOOKS APP !!\n")
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Println("3. Lihat Koleksi Buku")
		fmt.Println("10. List User")
		fmt.Println("99. Keluar\n")
		fmt.Print("Pilih Menu : ")
		fmt.Scanln(&input)

		switch input {
		case 1:
			var newUser entity.User
			fmt.Print("\nMasukan Nama: ")
			fmt.Scanln(&newUser.Nama)
			fmt.Print("Masukan No HP: ")
			fmt.Scanln(&newUser.No_hp)
			fmt.Print("Masukan Email: ")
			fmt.Scanln(&newUser.Email)
			fmt.Print("Masukan Password: ")
			fmt.Scanln(&newUser.Password)
			res := aksesUser.RegisUser(newUser)
			if res.ID == 0 {
				fmt.Println("Registrasi Gagal")
				break
			}
			fmt.Println("Registrasi Berhasil\n")

		case 2:
			var eemail string
			var pass string
			fmt.Println("\n--Login--")
			fmt.Print("Email: ")
			fmt.Scanln(&eemail)
			fmt.Print("Password: ")
			fmt.Scanln(&pass)
			emailauth := aksesUser.GetUserEmail(eemail)
			passauth := aksesUser.GetUserPass(eemail, pass)
			if emailauth == true && passauth == true {
				fmt.Println("Login Berhasil")
			} else if emailauth == false || passauth == false {
				fmt.Println("Email dan Password tidak sesuai, silahkan coba lagi")
			}

		case 10:
			fmt.Println("Daftar Seluruh User")
			for _, val := range aksesUser.GetAllData() {
				fmt.Println(val)
			}
		default:
			continue
		}
	}
	fmt.Println("TERIMA KASIH TELANG MENGGUNAKAN APLIKASI KAMI")
}
