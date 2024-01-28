package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(2)

	// Penerapan Buffer Channel
	// messages := make(chan int, 3)
	// go func() {
	// 	for {
	// 		i := <-messages
	// 		fmt.Println("receive data", i)
	// 	}
	// }()

	// for i := 0; i < 5; i++ {
	// 	fmt.Println("send data", i)
	// 	messages <- i
	// }

	// Penerapan Select Channel
	// var numbers = []int{3, 4, 3, 5, 6, 3, 2, 2, 6, 3, 4, 6, 3}
	// fmt.Println("numbers :", numbers)
	// var ch1 = make(chan float64)
	// go getAverage(numbers, ch1)
	// var ch2 = make(chan int)
	// go getMax(numbers, ch2)

	// for i := 0; i < 2; i++ {
	// 	select {
	// 	case avg := <-ch1:
	// 		fmt.Printf("Avg \t: %.2f \n", avg)
	// 	case max := <-ch2:
	// 		fmt.Printf("Max \t: %d \n", max)
	// 	}
	// }

	// Penerapan Mutex
	var wg sync.WaitGroup
	var meter counter
	var mtx sync.Mutex
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 1000; j++ {
				mtx.Lock()
				meter.Add(1)
				mtx.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(meter.val)

}

// Penerapan Select Channel
// func getAverage(numbers []int, ch chan float64) {
// 	var sum = 0
// 	for _, e := range numbers {
// 		sum += e
// 	}

// 	ch <- float64(sum) / float64(len(numbers))
// }

// func getMax(numbers []int, ch chan int) {
// 	var max = numbers[0]
// 	for _, e := range numbers {
// 		if max < e {
// 			max = e
// 		}
// 	}
// 	ch <- max
// }

// Penerapan Mutex
type counter struct {
	sync.Mutex
	val int
}

func (c *counter) Add(x int) {
	c.Lock()
	c.val++
	c.Unlock()
}
