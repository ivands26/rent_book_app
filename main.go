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

func deleteBook() int {
	var inputBook int
	fmt.Println("\nMasukkan ID buku yang ingin anda hapus")
	fmt.Scanln(&inputBook)
	return inputBook
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
					fmt.Println("4. Add My Book")
					fmt.Println("5. Edit My Book")
					fmt.Println("6. Delete My Book")
					fmt.Println("7. Rent Book") //harus input id user dan id book
					fmt.Println("8. Rent Book List")
					fmt.Println("9. Return Book")
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
					case 2:
						var input3 int = 0
						for input3 != 33 {
							fmt.Println("\nChoose what you want to edit: ")
							fmt.Println("1. Nama")
							fmt.Println("2. No HP")
							fmt.Println("3. Email")
							fmt.Println("4. Password")
							fmt.Println("33. Back to previous page\n")
							fmt.Print("Choose Menu : ")
							fmt.Scan(&input3)
							if input3 == 33 {
								input2 = 55
							}

							switch input3 {

							case 1:
								var namaUpdate string
								fmt.Println("\n-----Update Nama-----")
								for _, value := range aksesUser.GetProfileUser(email) {
									fmt.Println("Current Name : ", value.Nama)
								}
								fmt.Print("New Name : ")
								fmt.Scanln(&email)
								fmt.Scanln(&namaUpdate)
								res := aksesUser.UpdateUserNama(email, namaUpdate)
								if res == true {
									fmt.Println("Update Succes")
								} else {
									fmt.Println("Update Failed, Try Again")
								}
							case 2:
								var hpUpdate string
								fmt.Println("\n-----Update No. HP-----")
								for _, value := range aksesUser.GetProfileUser(email) {
									fmt.Println("Current Phone Number : ", value.No_hp)
								}
								fmt.Print("New Phone Number : ")
								fmt.Scanln(&email)
								fmt.Scanln(&hpUpdate)
								res := aksesUser.UpdateUserNo(email, hpUpdate)
								if res == true {
									fmt.Println("Update Succes")
								} else {
									fmt.Println("Update Failed, Try Again")
								}
							case 3:
								var surelUpdate string
								fmt.Println("\n-----Update Email-----")
								for _, value := range aksesUser.GetProfileUser(email) {
									fmt.Println("Current Email : ", value.Email)
								}
								fmt.Print("New Email : ")
								fmt.Scanln(&email)
								fmt.Scanln(&surelUpdate)
								res := aksesUser.UpdateUserSurel(email, surelUpdate)
								if res == true {
									fmt.Println("Update Succes")
								} else {
									fmt.Println("Update Failed, Try Again")
								}
							case 4:
								var PassUpdate string
								fmt.Println("\n-----Update Password-----")
								for _, value := range aksesUser.GetProfileUser(email) {
									fmt.Println("Current Password : ", value.Password)
								}
								fmt.Print("New Password : ")
								fmt.Scanln(&email)
								fmt.Scanln(&PassUpdate)
								res := aksesUser.UpdateUserPass(email, PassUpdate)

								if res == true {
									fmt.Println("Update Succes")
								} else {
									fmt.Println("Update Failed, Try Again")
								}
							default:
								continue
							}

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
						fmt.Print("\nMasukkan Book Title: ")
						fmt.Scanln(&newBook.Judul)
						fmt.Print("Masukkan Nama Author : ")
						fmt.Scanln(&newBook.Author)
						for _, val := range aksesUser.GetProfileUser(email) {
							newBook.Owned_by = val.Nama
						}
						aksesBook.InputBook(newBook)
						if res.ID == 0 {
							fmt.Println("Buku Gagal Diinput")
							break
						}
						fmt.Println("Buku Berhasil Diinput\n")

					case 6:
						inputBook := deleteBook()
						fmt.Println(aksesBook.DeleteBook(inputBook))
					default:
						continue
					}

				}
			}

		case 3:
			fmt.Println("\nDaftar Buku")
			for _, value := range aksesBook.GetDataBook() {

				fmt.Print("\nTitle\t: ")
				fmt.Println(value.Judul)
				fmt.Print("Isbn\t: ")
				fmt.Println(value.ID)
				fmt.Print("Author\t: ")
				fmt.Println(value.Author)
				fmt.Print("Owned by: ")
				fmt.Println(value.Owned_by)
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
