package main

import "fmt"

func main() {
	// TYPE CASTING
	var a float64 = float64(24)
	fmt.Println(a) // 24

	var b int32 = int32(24.00)
	fmt.Println(b) // 24


	// TYPE INFERENCE
	var firstName string = "my"
	lastName := "skill"
	fmt.Printf("halo %s %s\n", firstName, lastName)


	// TYPE ARRAY
	var fruits = [4]string{"apple","grape","banana","melon"}
	fmt.Println("Jumlah element \t\t", len(fruits)) 
	fmt.Println("Isi semua element \t", fruits)
	fmt.Println("Index kesatu \t", fruits[0])

	// TYPE MAPS
	var chicken map[string]int
	chicken = map[string]int{}
	chicken["januari"] = 50
	chicken["februari"] = 40
	fmt.Println("januari",chicken["januari"]) // januari 50 
	fmt.Println("mei", chicken["mei"]) // mei 0

	// TYPE STRUCTS
	type student struct {
		name string
		grade int
	} 

	var s1 student
	s1.name = "my skill"
	s1.grade = 2
	fmt.Println("name :", s1.name)
	fmt.Println("grade :", s1.grade)

	// MAKE
	var fruitsX = make([]string, 2)
	fruitsX[0] = "Jeruk"
	fruitsX[1] = "Durian"
	fmt.Println(fruitsX) // [Jeruk Durian]
}