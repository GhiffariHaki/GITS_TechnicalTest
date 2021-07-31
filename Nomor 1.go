package main

import (
	"fmt"
)

//Fungsi untuk mengecek apakah bilangan dapat dibagi habis oleh 3 atau 5
func modulo(angka int) {
	if (angka%3 == 0) && (angka%5 == 0) {
		//Kondisi bilangan dapat dibagi habis oleh 3 dan 5
		fmt.Println("Hello World")
	} else if angka%3 == 0 {
		//Kondisi bilangan dapat dibagi habis oleh 3
		fmt.Println("Hello")
	} else if angka%5 == 0 {
		//Kondisi bilangan dapat dibagi habis oleh 5
		fmt.Println("World")
	} else {
		//Kondisi bilangan dapat dibagi habis oleh 3 atau 5
		fmt.Println("Angka tidak bisa dibagi habis oleh 3 atau 5")
	}
}

// Fungsi Utama
func main() {
	var angka, test int

	fmt.Println("Masukkan Test Case :")
	fmt.Scanln(&test)
	for i := 0; i < test; i++ {
		//Meminta Input Angka
		fmt.Println("Masukkan Angka :")
		fmt.Scanln(&angka)
		//Memasukkan Angka yang diinput kedalam fungsi modulo
		modulo(angka)
	}
	//Pemberitahuan bahwa program telah berakhir
	fmt.Println("Program Berakhir.")
}
