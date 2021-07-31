package main

import (
	"fmt"
)

//Fungsi untuk membalik kata
func reverse(kata string) {
	//Variabel untuk menyimpan hasil
	var hasil string
	for _, s := range kata {
		//Hasil string dibentuk dari hasil append pecahan-pecahan string kata (s)
		hasil = string(s) + hasil
	}
	//Print Hasil
	fmt.Println(hasil)

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
		//Memasukkan kata yang diinput kedalam fungsi reverse
		reverse(kata)
	}
	//Pemberitahuan bahwa program telah berakhir
	fmt.Println("Program Berakhir.")
}
