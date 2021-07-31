package main

import (
	"fmt"
	"strings"
)

//Fungsi untuk membalik kata
func reverse(kata string) (hasil string) {
	for _, s := range kata {
		//Hasil string dibentuk dari hasil append pecahan-pecahan string kata (s)
		hasil = string(s) + hasil
	}
	return
}

//Fungsi untuk mengecek palindrom
func palindrom(kata string) {
	var katarev, katarevlow, katalow string
	//Membuat reverse dari kata yang diinput
	katarev = reverse(kata)
	//Mengubah kedua string menjadi lowercase only
	katarevlow = strings.ToLower(katarev)
	katalow = strings.ToLower(kata)
	//Mengecek apakah kedua string sama persis
	result := katalow == katarevlow
	//Print Hasil
	fmt.Println("Hasil : ", result)
}

// Fungsi Utama
func main() {
	var test int
	var kata string

	fmt.Println("Masukkan Test Case :")
	fmt.Scanln(&test)
	for i := 0; i < test; i++ {
		//Meminta Input Kata
		fmt.Println("Masukkan Kata :")
		fmt.Scanln(&kata)
		//Memasukkan Angka yang diinput kedalam fungsi modulo
		palindrom(kata)
	}
	//Pemberitahuan bahwa program telah berakhir
	fmt.Println("Program Berakhir.")
}
