package main

import (
	"errors"
	"fmt"
	"strings"
)

// "os"

func main() {
	// Penerapan defer
	// defer fmt.Println("halo")
	// fmt.Println("selamat datang")

	// Penerapan exit
	// fmt.Println("fitria")
	// os.Exit(1)
	// fmt.Println("melliyanni")

	// Penerapan Error
	// var input string
	// fmt.Print("Type some number: ")
	// fmt.Scanln(&input)
	// var number int
	// var err error
	// number, err = strconv.Atoi(input)
	// if err == nil {
	// 	fmt.Println(number, "is number")
	// } else {
	// 	fmt.Println(input, "is not number")
	// 	fmt.Println(err.Error())
	// }

	// Penerapan Custom Error
	// var name string
	// fmt.Print("Type your name: ")
	// fmt.Scanln(&name)
	// if valid, err := validate(name);
	// valid {
	// 	fmt.Println("halo", name)
	// } else {
	// 	fmt.Println(err.Error())
	// }

	// Penerapan Panic
	// var name string
	// fmt.Print("Type your name:")
	// fmt.Scanln(&name)
	// if valid, err := validate(name);
	// valid {
	// 	fmt.Println("halo", name)
	// } else {
	// 	defer fmt.Println("continue")
	// 	panic(err.Error())
	// 	fmt.Println("end")
	// }

	// Penerapan Recover
	defer catch()
	var name string
	fmt.Print("Type your name: ")
	fmt.Scanln(&name)
	if valid, err := validate(name);
	valid{
		fmt.Println("halo", name)
	} else {
		panic(err.Error())
		fmt.Println("end")
	}
}

// Penerapan Custom Error
func validate (input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("cannot be empty")
	}
	return true, nil
}


// Penerapan Recover
func catch() {
	if r := recover();
	r != nil {
		fmt.Println("Error occured", r)
	} else {
		fmt.Println("Application running perfectly")
	}
}