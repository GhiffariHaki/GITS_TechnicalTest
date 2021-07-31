package main

import (
	"fmt"
	"strings"
)

//Fungsi untuk mengecek apakah email sesuai
func check(email string) {
	//Mengecek apakah string yang diinput memiliki simbol @
	if strings.Contains(email, "@") {
		//Memisahkan string dengan simbol @ lalu hasilnya di assign kedalam variabel
		s := strings.Split(email, "@")
		awal, akhir := s[0], s[1]
		//fmt.Println(awal)
		//fmt.Println(akhir)
		if len(awal) > 20 {
			//Mengecek apakah string sebelum tanda @ lebih dari 20 character
			fmt.Println("False, Nama email maksimum 20 Character")
		} else if strings.Contains(akhir, ".") {
			//Mengecek apakah domain mengandung co.id atau .id
			if strings.Contains(akhir, ".co.id") {
				fmt.Println("True")
			} else if strings.Contains(akhir, ".id") {
				fmt.Println("True")
			} else {
				//Email yang diinput tidak co.id atau .id
				fmt.Println("False, Domain yang diperbolehkan hanya id atau co.id")
			}
		} else {
			//Email yang diinput tidak memiliki "." setelah simbol @
			fmt.Println("False, harus ada titik setelah @")
		}

	} else {
		//Email tidak mengandung @
		fmt.Println("False, Email harus mengandung @")
	}
}

// Fungsi Utama
func main() {
	var test int
	var email string

	fmt.Println("Masukkan Test Case :")
	fmt.Scanln(&test)
	for i := 0; i < test; i++ {
		//Meminta Input Email
		fmt.Println("Masukkan Kata :")
		fmt.Scanln(&email)
		//Memasukkan email yang diinput kedalam fungsi check email
		check(email)
	}
	//Pemberitahuan bahwa program telah berakhir
	fmt.Println("Program Berakhir.")
}
