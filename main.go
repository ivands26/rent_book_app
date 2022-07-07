package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"group_project/config"
	"group_project/entity"
)

func halamanregis() entity.User {
	var newUser entity.User
	in := bufio.NewReader(os.Stdin)
	fmt.Print("Masukan Nama: ")
	newUser.Nama, _ = in.ReadString('\n')
	newUser.Nama = strings.TrimSpace(newUser.Nama)
	for newUser.Nama == "" {
		fmt.Print("\nNama Tidak Boleh Kosong !\n\n")
		fmt.Print("\nMasukan Nama Sekali Lagi: ")
		newUser.Nama, _ = in.ReadString('\n')
		newUser.Nama = strings.TrimSpace(newUser.Nama)
	}
	fmt.Print("Masukan No HP: ")
	fmt.Scanln(&newUser.No_hp)
	for newUser.No_hp == "" {
		fmt.Print("\nNo HP Tidak Boleh Kosong !\n\n")
		fmt.Print("\nMasukan No HP Sekali Lagi: ")
		fmt.Scanln(&newUser.No_hp)
	}
	fmt.Print("Masukan Email: ")
	fmt.Scanln(&newUser.Email)
	for newUser.Email == "" {
		fmt.Print("\nEmail Tidak Boleh Kosong !\n\n")
		fmt.Print("\nMasukan Email Sekali Lagi: ")
		fmt.Scanln(&newUser.Email)
	}
	fmt.Print("Masukan Password: ")
	fmt.Scanln(&newUser.Password)
	for newUser.Password == "" {
		fmt.Print("\nPassword Tidak Boleh Kosong !\n\n")
		fmt.Print("\nMasukan Password Sekali Lagi: ")
		fmt.Scanln(&newUser.No_hp)
	}
	return newUser
}
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

func validasiemailpass(emailauth, passauth, emailpassauth bool) bool {
	var has bool
	if !emailauth && !passauth {
		fmt.Println("Email dan Password tidak sesuai, silahkan coba lagi")
		has = false
	} else if !passauth {
		fmt.Println("Password salah")
		has = false
	} else if !emailauth {
		fmt.Println("Email tidak terdaftar")
		has = false
	} else if emailpassauth == true {
		fmt.Println("Login Berhasil")
		has = true
	}
	return has
}

func halamanmyprofile(email string) {
	conn := config.InitDB()
	aksesUser := entity.AksesUser{DB: conn}
	Val := aksesUser.GetProfileUser(email)
	fmt.Println("\nProfile")
	fmt.Print("ID\t: ")
	fmt.Println(Val.ID)
	fmt.Print("Nama\t: ")
	fmt.Println(Val.Nama)
	fmt.Print("No HP\t: ")
	fmt.Println(Val.No_hp)
	fmt.Print("Email\t: ")
	fmt.Println(Val.Email)
}

func halamaneditprofile() {
	fmt.Println("\nChoose what you want to edit: ")
	fmt.Println("1. Nama")
	fmt.Println("2. No HP")
	fmt.Println("3. Email")
	fmt.Println("4. Password")
	fmt.Print("33. Back to previous page\n\n")
}

func halamaneditprofile2(email string, input2 int, input3 int) (int, int) {
	conn := config.InitDB()
	aksesUser := entity.AksesUser{DB: conn}
	Val := aksesUser.GetProfileUser(email)
	switch input3 {
	case 1:
		var namaUpdate string
		in := bufio.NewReader(os.Stdin)
		fmt.Println("\n-----Update Nama-----")
		fmt.Println("Current Name : ", Val.Nama)
		fmt.Print("New Name : ")
		fmt.Scanln(&email)
		namaUpdate, _ = in.ReadString('\n')
		namaUpdate = strings.TrimSpace(namaUpdate)
		res := aksesUser.UpdateUserNama(email, namaUpdate)
		if res == true {
			fmt.Println("Update Succes")
		} else {
			fmt.Println("Update Failed, Try Again")
		}

	case 2:
		var hpUpdate string
		fmt.Println("\n-----Update No. HP-----")
		fmt.Println("Current Phone Number : ", Val.No_hp)
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
		fmt.Println("Current Email : ", Val.Email)
		fmt.Print("New Email : ")
		fmt.Scanln(&email)
		fmt.Scanln(&surelUpdate)
		res := aksesUser.UpdateUserSurel(email, surelUpdate)
		if res == true {
			fmt.Print("Update Berhasil, anda akan kembali ke halaman utama, silahkan login kembali\n\n")
			input3 = 33
			input2 = 80
			break
		} else {
			fmt.Println("Update Failed, Try Again")
		}
	case 4:
		var PassUpdate string
		fmt.Println("\n-----Update Password-----")
		fmt.Println("Current Password : ", Val.Password)
		fmt.Print("New Password : ")
		fmt.Scanln(&email)
		fmt.Scanln(&PassUpdate)
		res := aksesUser.UpdateUserPass(email, PassUpdate)
		if res == true {
			fmt.Println("Update Succes")
		} else {
			fmt.Println("Update Failed, Try Again")
		}
	}
	return input3, input2
}
func halamandeleteaccount() int {
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
	fmt.Print("\nMasukkan ID buku yang ingin anda hapus : ")
	fmt.Scanln(&inputBook)
	return inputBook
}

func halamanlistmybook(email string) {
	conn := config.InitDB()
	aksesUser := entity.AksesUser{DB: conn}
	aksesBook := entity.AksesBook{DB: conn}
	Val := aksesUser.GetProfileUser(email)
	fmt.Print("\n------List My Book-----\n\n")
	for _, val := range aksesBook.GetMyBook(email) {
		fmt.Print("\nID Book\t: ")
		fmt.Println(val.ID)
		fmt.Print("Title\t: ")
		fmt.Println(val.Judul)
		fmt.Print("Isbn\t: ")
		fmt.Println(val.ISBN)
		fmt.Print("Author\t: ")
		fmt.Println(val.Author)
		fmt.Print("Owned by: ")
		fmt.Println(val.Nama)
		fmt.Print("Status\t: ")
		fmt.Println(val.Status)
	}
	var IDUser uint = Val.ID
	fmt.Print("\n\nList My Rented Book\n\n")
	if len(aksesBook.GetDataYourRentedBook(IDUser)) == 0 {
		fmt.Println("---No book Rented---")
	} else {
		for _, val := range aksesBook.GetDataYourRentedBook(IDUser) {
			fmt.Print("\nID Book\t: ")
			fmt.Println(val.ID)
			fmt.Print("Title\t: ")
			fmt.Println(val.Judul)
			fmt.Print("Isbn\t: ")
			fmt.Println(val.ISBN)
			fmt.Print("Author\t: ")
			fmt.Println(val.Author)

			ownerid := val.Owned_by
			own := aksesUser.GetOwner(ownerid)
			fmt.Print("Owned by: ")
			fmt.Println(own.Nama)
		}
	}
}

func halamanrentbook(email string) {
	conn := config.InitDB()
	aksesUser := entity.AksesUser{DB: conn}
	aksesBook := entity.AksesBook{DB: conn}
	aksesRent := entity.AksesRent{DB: conn}
	Val := aksesUser.GetProfileUser(email)
	var bookstatus string = "Available"
	fmt.Print("\n------Rent Book------\n\n")
	fmt.Print("Available Book For Rent\n\n")
	ab := aksesBook.GetDataRentBook(email, bookstatus)
	if len(ab) == 0 {
		fmt.Println("---No Book Available For Rent---")
	} else {
		for _, val := range ab {
			fmt.Print("\nID Book\t: ")
			fmt.Println(val.ID)
			fmt.Print("Title\t: ")
			fmt.Println(val.Judul)
			fmt.Print("Isbn\t: ")
			fmt.Println(val.ISBN)
			fmt.Print("Author\t: ")
			fmt.Println(val.Author)
			fmt.Print("Owned by: ")
			fmt.Println(val.Nama)
		}
		var IDBuku uint
		var IDUser uint = Val.ID
		var stat string = "Rented"
		fmt.Print("\nMasukkan ID Book yang ingin dipinjam : ")
		fmt.Scanln(&IDBuku)
		res := aksesRent.RentBuku(IDUser, IDBuku)
		status := aksesBook.Updatestatus(IDBuku, stat)
		if !res && !status {
			fmt.Println("Gagal untuk meminjam, silahkan coba lagi")
		} else {
			fmt.Print("\nBerhasil meminjam buku. Durasi peminjaman adalah 7 hari, silahkan kembalikan tepat waktu !\n")
		}
	}
}

func halamanreturnbook(email string) {
	conn := config.InitDB()
	aksesUser := entity.AksesUser{DB: conn}
	aksesBook := entity.AksesBook{DB: conn}
	aksesRent := entity.AksesRent{DB: conn}
	Val := aksesUser.GetProfileUser(email)
	var IDUser uint = Val.ID
	fmt.Print("\n-----Return Book-----\n\n")
	fmt.Print("List Your Rented Book\n\n")
	if len(aksesBook.GetDataYourRentedBook(IDUser)) == 0 {
		fmt.Println("---No book Rented---")
	} else {
		for _, val := range aksesBook.GetDataYourRentedBook(IDUser) {
			fmt.Print("\nID Book\t: ")
			fmt.Println(val.ID)
			fmt.Print("Title\t: ")
			fmt.Println(val.Judul)
			fmt.Print("Isbn\t: ")
			fmt.Println(val.ISBN)
			fmt.Print("Author\t: ")
			fmt.Println(val.Author)
			fmt.Print("Owned by: ")
			fmt.Println(val.Owned_by)
		}
		var IDBuku uint
		var stat string = "Available"
		fmt.Print("\nMasukkan ID Book yang dikembalikan : ")
		fmt.Scanln(&IDBuku)
		res := aksesRent.ReturnBuku(IDBuku)
		status := aksesBook.Updatestatus(IDBuku, stat)
		if !res && !status {
			fmt.Println("Gagal untuk kembalikan buku, silahkan coba lagi")
		}
		fmt.Print("\nBerhasil kembalikan buku. Terimakasih sudah mengembalikan buku tepat waktu!\n")
	}

}

func main() {
	conn := config.InitDB()
	config.MigrateDB(conn)
	aksesUser := entity.AksesUser{DB: conn}
	aksesBook := entity.AksesBook{DB: conn}

	var email, pass string
	var input int
	var input2 int
	for input != 99 {
		fmt.Println("\n\t\t==================================")
		fmt.Println("\t\t|| Welcome in RENT BOOKS APP !! ||")
		fmt.Print("\t\t==================================\n\n")
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Println("3. Lihat Koleksi Buku")
		fmt.Println("10. List User")
		fmt.Println("99. Keluar")
		fmt.Print("\nPilih Menu : ")
		fmt.Scanln(&input)

		switch input {
		case 1:
			newUser := halamanregis()
			res := aksesUser.RegisUser(newUser)
			if res.ID == 0 {
				fmt.Println("Registrasi Gagal")
				break
			}
			fmt.Print("\n--> REGISTRASI BERHASIL <--\n")

		case 2:
			email, pass = halamanlogin()
			emailauth := aksesUser.GetUserEmail(email)
			passauth := aksesUser.GetUserPass(pass)
			emailpassauth := aksesUser.GetEmailPass(email, pass)
			cek := validasiemailpass(emailauth, passauth, emailpassauth)
			if cek == true {
				input2 = 0
				for input2 != 80 {
					Val := aksesUser.GetProfileUser(email)
					fmt.Printf("\n\t---Welcome %s !!---\n\n", Val.Nama)
					fmt.Println("1. My Profile")
					fmt.Println("2. Edit Profile")
					fmt.Println("3. Delete Account")
					fmt.Println("4. Add My Book")
					fmt.Println("5. List My Book and My Rented Book")
					fmt.Println("6. Edit My Book")
					fmt.Println("7. Delete My Book")
					fmt.Println("8. Rent Book")
					fmt.Println("9. Return Book")
					fmt.Print("80. Log Out \n\n")
					fmt.Print("Pilih Menu : ")
					fmt.Scanln(&input2)

					switch input2 {
					case 1:
						halamanmyprofile(email)
						fmt.Print("\n55. Kembali ke menu sebelumnya\n\n")
						fmt.Print("Pilih Menu : ")
						fmt.Scanln(&input2)

					case 2:
						var input3 int = 0
						for input3 != 33 {
							halamaneditprofile()
							fmt.Print("Choose Menu : ")
							fmt.Scan(&input3)
							if input3 == 33 {
								input2 = 55
							}
							input3, input2 = halamaneditprofile2(email, input2, input3)
						}
					case 3:
						inputYT := halamandeleteaccount()
						if inputYT == 1 {
							aksesUser.DeleteUser(email)
							fmt.Println("\nAKUN BERHASIL DIHAPUS")
							input2 = 80
						} else {
							input2 = inputYT
						}

					case 4:
						in := bufio.NewReader(os.Stdin)
						var newBook entity.Book
						fmt.Print("\nMasukkan Book Title: ")
						newBook.Judul, _ = in.ReadString('\n')
						newBook.Judul = strings.TrimSpace(newBook.Judul)
						fmt.Print("Masukkan Nama Author : ")
						newBook.Author, _ = in.ReadString('\n')
						newBook.Author = strings.TrimSpace(newBook.Author)
						newBook.Owned_by = Val.ID
						newBook.Status = "Available"
						res := aksesBook.InputBook(newBook)
						if res.ID == 0 {
							fmt.Println("Buku Gagal Diinput")
							break
						}
						fmt.Println("Buku Berhasil Diinput")
					case 5:
						halamanlistmybook(email)
						fmt.Print("\n55. Kembali ke menu sebelumnya\n\n")
						fmt.Print("Pilih Menu : ")
						fmt.Scanln(&input2)
					case 6:
						var judulUpdate string
						var authorUpdate string
						var id int
						fmt.Println("\n-----Update Buku-----")
						for _, val := range aksesBook.GetMyBook(email) {
							fmt.Print("\nID Book\t: ")
							fmt.Println(val.ID)
							fmt.Print("Title\t: ")
							fmt.Println(val.Judul)
							fmt.Print("Isbn\t: ")
							fmt.Println(val.ISBN)
							fmt.Print("Author\t: ")
							fmt.Println(val.Author)
						}
						in := bufio.NewReader(os.Stdin)
						fmt.Print("\nMasukkan ID Buku : ")
						fmt.Scanln(&id)
						fmt.Print("New Judul : ")
						judulUpdate, _ = in.ReadString('\n')
						judulUpdate = strings.TrimSpace(judulUpdate)
						for judulUpdate == "" {
							fmt.Print("\nNama Tidak Boleh Kosong !\n\n")
							fmt.Print("\nMasukan Nama Sekali Lagi: ")
							judulUpdate, _ = in.ReadString('\n')
							judulUpdate = strings.TrimSpace(judulUpdate)
						}
						fmt.Print("New Author : ")
						authorUpdate, _ = in.ReadString('\n')
						authorUpdate = strings.TrimSpace(authorUpdate)
						for authorUpdate == "" {
							fmt.Print("\nNama Tidak Boleh Kosong !\n\n")
							fmt.Print("\nMasukan Nama Sekali Lagi: ")
							authorUpdate, _ = in.ReadString('\n')
							authorUpdate = strings.TrimSpace(authorUpdate)
						}
						res := aksesBook.UpdateBookJA(id, judulUpdate, authorUpdate)
						if !res {
							fmt.Println("Update Failed, Try Again")
						} else {
							fmt.Println("Update Succes")
						}

					case 7:
						fmt.Println("\n-----Delete Buku-----")
						for _, val := range aksesBook.GetMyBook(email) {
							fmt.Print("\nID Book\t: ")
							fmt.Println(val.ID)
							fmt.Print("Title\t: ")
							fmt.Println(val.Judul)
							fmt.Print("Isbn\t: ")
							fmt.Println(val.ISBN)
							fmt.Print("Author\t: ")
							fmt.Println(val.Author)
						}
						inputBook := deleteBook()
						res := aksesBook.DeleteBook(inputBook)
						if !res {
							fmt.Println("Delete Failed, Try Again")
						} else {
							fmt.Println("Delete Success")
						}
					case 8:
						halamanrentbook(email)

					case 9:
						halamanreturnbook(email)

					default:
						continue
					}

				}
			}
		case 3:
			fmt.Println("\nDaftar Buku")
			for _, value := range aksesBook.GetDataBook() {
				fmt.Print("\nID Book\t: ")
				fmt.Println(value.ID)
				fmt.Print("Title\t: ")
				fmt.Println(value.Judul)
				fmt.Print("Isbn\t: ")
				fmt.Println(value.ISBN)
				fmt.Print("Author\t: ")
				fmt.Println(value.Author)
				fmt.Print("Owned by: ")
				fmt.Println(value.Nama)
				fmt.Print("Status\t: ")
				fmt.Println(value.Status)
			}
		default:
			continue
		}
	}
	fmt.Println("\nTERIMA KASIH TELAH MENGGUNAKAN APLIKASI KAMI :)")
}
