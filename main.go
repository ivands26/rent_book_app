package main

import (
	"fmt"

	"group_project/config"
	"group_project/entity"
)

func halamanlogin() (string, string) {
	var email string
	var pass string
	fmt.Println("\n--Login--")
	fmt.Print("Email: ")
	fmt.Scanln(&email)
	fmt.Print("Password: ")
	fmt.Scanln(&pass)
	return email, pass
}

func deleteAccount() int {
	var inputYT int
	fmt.Println("\nApakah anda yakin ingin menghapus akun anda ?")
	fmt.Println("1. Ya dan kembali ke halaman utama")
	fmt.Println("55. Tidak dan kembali ke menu sebelumnya")
	fmt.Print("\nPilih Menu : ")
	fmt.Scanln(&inputYT)
	return inputYT
}
func main() {
	conn := config.InitDB()
	config.MigrateDB(conn)
	aksesUser := entity.AksesUser{DB: conn}
	aksesBook := entity.AksesBook{DB: conn}

	var input int
	var input2 int
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
			email, pass := halamanlogin()
			emailauth := aksesUser.GetUserEmail(email)
			passauth := aksesUser.GetUserPass(pass)
			emailpassauth := aksesUser.GetEmailPass(email, pass)
			if emailauth == false && passauth == false {
				fmt.Println("Email dan Password tidak sesuai, silahkan coba lagi")
			} else if passauth == false {
				fmt.Println("Password salah")
			} else if emailauth == false {
				fmt.Println("Email tidak terdaftar")
			} else if emailpassauth == true {
				fmt.Println("Login Berhasil")
				input2 = 0
				for input2 != 80 {
					for _, val := range aksesUser.GetName(email) {
						fmt.Printf("\n\tWelcome %s !!\n\n", val.Nama)
					}

					fmt.Println("1. My Profile")
					fmt.Println("2. Edit Profile")
					fmt.Println("3. Delete Account")
					fmt.Println("4. Tambahkan Buku")
					fmt.Println("5. Lihat Buku Saya")
					fmt.Println("6. Pinjam Buku") //nanti masukin list
					fmt.Println("7. Kembalikan Buku")
					fmt.Println("80. Log Out\n")
					fmt.Print("Pilih Menu : ")
					fmt.Scanln(&input2)

					switch input2 {
					case 1:
						fmt.Println("\nProfile")
						for _, val := range aksesUser.GetProfileUser(email) {
							fmt.Print("ID\t: ")
							fmt.Println(val.ID)
							fmt.Print("Nama\t: ")
							fmt.Println(val.Nama)
							fmt.Print("No HP\t: ")
							fmt.Println(val.No_hp)
							fmt.Print("Email\t: ")
							fmt.Println(val.Email)

							fmt.Println("\n55. Kembali ke menu sebelumnya \n")
							fmt.Print("Pilih Menu : ")
							fmt.Scanln(&input2)
						}
					case 3:
						inputYT := deleteAccount()
						if inputYT == 1 {
							aksesUser.DeleteUser(email)
							fmt.Println("\nAKUN BERHASIL DIHAPUS\n")
							input2 = 80
						} else {
							input2 = inputYT
						}

					case 4:
						var newBook entity.Book
						fmt.Print("Book Title: ")
						fmt.Scanln(&newBook.Judul)
						fmt.Print("Author : ")
						fmt.Scanln(&newBook.Author)
						res := aksesBook.InputBook(newBook)
						if res.ID == 0 {
							fmt.Println("Buku Gagal Diinput")
							break
						}
						fmt.Println("Buku Berhasil Diinput\n")

					default:
						continue
					}

				}
			}

		case 3:
			fmt.Println("Daftar Buku")
			for _, value := range aksesBook.GetDataBook() {
				fmt.Println(value.Judul)
				fmt.Println(value.Author)
				fmt.Print("\n")
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
	fmt.Println("\nTERIMA KASIH TELAH MENGGUNAKAN APLIKASI KAMI :)")
}
