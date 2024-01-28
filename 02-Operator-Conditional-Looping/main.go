package main

import "fmt"

func main()  {
	// OPERATOR ARITMATIKA
	var value = (((2+6)%3)*4-2)/3

	var valModulus = 10 % 6
	fmt.Println(value)
	fmt.Println(valModulus)

	// OPERATOR RELASIONAL (PERBANDINGAN)
	var valRelasional = (((2+3)%3)*4-2) / 3
	var isEqual = (valRelasional == 2)
	fmt.Printf("nilai %d (%t) \n", valRelasional, isEqual)

	// OPERATOR LOGIKA
	var valLogic = false && true
	fmt.Println(valLogic)

	// KONDISIONAL
	var point = 2
	if point == 10 {
		fmt.Println("Passed with the perfect score")
	} else if point > 5 {
		fmt.Println("Passed")
	} else if point == 4 {
		fmt.Println("Almost passed")
	} else {
		fmt.Println("Failed passed, your point is", point)
	}

	// SWITCH
	var pointB = 7

	switch pointB{
	case 8:
		fmt.Println("Perfect")
	case 7:
		fmt.Println("Awesome")
	default:
		fmt.Println("Not Bad")
	}

	// SWITCH FALLTHROUGH
	var pointC = 14

	switch {
	case pointC == 8:
		fmt.Println("Perfect")
	case (pointC < 8) && (pointC > 3):
		fmt.Println("Awesome")
		fallthrough
	case pointC < 5:
		fmt.Println("You need to learn more")
	default:
		{
			fmt.Println("Not Bad")
			fmt.Println("You need to learn more")
		}
	}

	// PERULANGAN
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}

	// PERULANGAN BERSARANG
	for y := 0; y < 5; y++ {
		for j :=y; j < 5; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
}