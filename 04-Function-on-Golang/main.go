package main

import (
	"fmt"
	"math"
	"strings"
)

// FUNCTION
func main() {
	var names = []string{"my", "skill"}
	
	printMessage("halo", names)

	area, circum := calculate(3.2)
	fmt.Println(area)
	fmt.Println(circum)

	var avg = calcVariadic(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
	var avgMessage = fmt.Sprintf("Rata-rata : %.2f", avg)
	fmt.Println(avgMessage)

	// contoh function closure
		var getMinMax = func (n []int) (int, int)  {
			var min, max int
			for i, e := range n {
				switch {
				case i == 0:
					max, min = e, e
				case e > max:
					max = e
				case e < min:
					min = e
				}
			}

			return min, max
		}

		var numbersYY = []int{2, 3, 4, 3, 4, 2, 3}
		var min, max = getMinMax(numbersYY)
		fmt.Printf("data : %v\nmin : %v\nmax : %v\n", numbersYY, min, max)
	// end contoh function closure

	var hasil = filter([]string{"ini", "data"}, func(each string) bool {
		return true
	})

	fmt.Println(hasil)
}

func printMessage(message string, arr []string){
	var nameString = strings.Join(arr, " ")
	fmt.Println(message, nameString)
}

// FUNCTION MULTIPLE RETURN

func calculate(d float64) (float64, float64) { // hitung luas
	var area = math.Pi * math.Pow(d / 2,2) // hitung keliling
	var circumference = math.Pi * d
	// kembalikan 2 nilai
	return area, circumference
}

// FUNCTION VARIADIC : FUNCTION YANG BISA PUNYA BANYAK PARAMETER TIDAK TERBATAS

func calcVariadic(numbers ...int) float64{
	var total int = 0
	for _, number := range numbers {
		total += number
	}

	var avg = float64(total) / float64(len(numbers))
	return avg
}


// FUNCTION CALLBACK
func filter(data []string, callback func(string) bool) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}

	return result
}