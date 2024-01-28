package main

import (
	"context"
	"fmt"
)

func main() {
	// runtime.GOMAXPROCS(2)

	// Penerapan Goroutine
	// go print(5, "halo")
	// go print(5, "apa kabar")
	// var input string
	// fmt.Scanln(&input)

	// Penerapan Channel
	// var messages = make(chan string)
	// var sayHelloTo = func(who string) {
	// 	var data = fmt.Sprintf("hello %s", who)
	// 	messages <- data
	// }
	// go sayHelloTo("john wick")
	// go sayHelloTo("ethan hunt")
	// go sayHelloTo("jason bourne")
	// var message1 = <-messages
	// fmt.Println(message1)
	// var message2 = <-messages
	// fmt.Println(message2)
	// var message3 = <-messages
	// fmt.Println(message3)

	// Penerapan Wait Group
	// var wg sync.WaitGroup
	// for i := 0; i < 5; i++ {
	// 	var data = fmt.Sprintf("data %d", i)
	// 	wg.Add(1)
	// 	go doPrint(&wg, data)
	// }

	// wg.Wait()

	// Penerapan Context
	contextParent := context.Background()
	ctx1 := context.WithValue(contextParent, "key1", "hello world")
	ctx2 := context.WithValue(ctx1, "key2", "hello girls")
	ctx3 := context.WithValue(ctx2, "key3", "hello boys")
	ctx4 := context.WithValue(contextParent, "key4", "hello children")
	ctx5 := context.WithValue(ctx1, "key5", "hello later")
	fmt.Println(ctx5.Value("key5"))
	fmt.Println(ctx5.Value("key4"))
	fmt.Println(ctx5.Value("key3"))
	fmt.Println(ctx5.Value("key2"))
	fmt.Println(ctx5.Value("key1"))
	fmt.Println(ctx3.Value("key1"))
	fmt.Println(ctx4.Value("key1"))
}

// Penerapan Goroutine
// func print(till int, message string) {
// 	for i := 0; i < till; i++ {
// 		fmt.Println((i + 1), message)
// 	}
// }

// Penerapan Wait Group
// func doPrint(wg *sync.WaitGroup, message string) {
// 	defer wg.Done()
// 	fmt.Println(message)
// }
