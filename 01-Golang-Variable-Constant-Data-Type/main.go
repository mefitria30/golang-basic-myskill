package main

import "fmt"

func main()  {
	// INI VARIABEL
	var firstName string = "Fitria"
	var lastName string = " Melliyanni"
	fmt.Print("halo ", firstName, lastName, "\n")

	// INI KONSTANTA
	const name string = "Fitria Melliyanni"
	fmt.Print("nice to meet you ", name, "\n")

	// TIPE DATA NUMERIK NON DESIMAL
	var bilanganBulat uint8 = 20
	fmt.Print("ini bilangan bulat ", bilanganBulat, "\n")

	// TIPE DATA NUMERIK DESIMAL
	var bilanganDesimal = 2.0
	fmt.Print("ini bilangan Desimal ", bilanganDesimal, "\n")

	// TIPE DATA BOOLEAN
	var varBool = true
	fmt.Print("ini Boolean ", varBool, "\n")

	// TIPE DATA STRING
	var message string = "Halo"
	fmt.Printf("message: %s \n", message)

	// TIPE DATA VARIABEL POINTER
	var numberA int = 4
	var numberB *int = &numberA // tanda ampersand digunakan untuk mengambil alamat memori
	fmt.Print("ini Pointer A ", numberA, "\n")
	fmt.Print("ini Pointer B ", numberB, "\n")
	fmt.Print("ini Pointer C ", *numberB, "\n") // tanda asterisk digunakan untuk mengambil nilai variabel
}