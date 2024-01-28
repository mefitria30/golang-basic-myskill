package main

import (
	"fmt"
	"math"
)

// digunakan untuk : Decode JSON Ke STRUCT && Decode Array JSON ke Array Objek
// type User struct {
// 	FullName string `json:"Name"`
// 	Age      int
// }

// Penerapan Interface
type hitung interface {
	luas() float64
	keliling() float64
}

type lingkaran struct {
	diameter float64
}

type persegi struct {
	sisi float64
}

func main() {
	// var jsonString = `{
	// 	"Name":"John Wick", 
	// 	"Age":27
	// }`

	// var jsonData = []byte(jsonString)

	// Decode JSON Ke STRUCT
	// var data User
	// var err = json.Unmarshal(jsonData, &data)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// fmt.Println("user:", data.FullName)
	// fmt.Println("age:", data.Age)


	// Decode JSON Ke INTERFACE
	// var data1 map[string]interface{}
	// json.Unmarshal(jsonData, &data1)
	// fmt.Println("user:", data1["Name"])
	// fmt.Println("age:", data1["Age"])
	
	// var data2 interface{}
	// json.Unmarshal(jsonData, &data2)
	// var decodedData = data2.(map[string]interface{})
	// fmt.Println("user:", decodedData["Name"])
	// fmt.Println("age:", decodedData["Age"])

	// Decode Array JSON ke Array Objek
	// var jsonString = `[
	// 	{
	// 		"Name":"john wick",
	// 		"Age":27
	// 	},
	// 	{
	// 		"Name":"ethan hunt",
	// 		"Age":32
	// 	}
	// ]`

	// var data []User
	// var err = json.Unmarshal([]byte(jsonString), &data)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// fmt.Println("User 1:", data[0].FullName)
	// fmt.Println("User 2:", data[1].FullName)

	// Encode Objeck ke JSON String
	// var object = []User{{"john wick", 27}, {"ethan hunt", 32}}

	// var jsonData, err = json.Marshal(object)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// var jsonString = string(jsonData)
	// fmt.Println(jsonString)

	// Penerapan Interface
	var bangunDatar hitung
	bangunDatar = persegi{10.0}
	fmt.Println("===== persegi")
	fmt.Println("luas :", bangunDatar.luas())
	fmt.Println("keliling :", bangunDatar.keliling())
	bangunDatar = lingkaran{14.0}
	fmt.Println("===== lingkaran")
	fmt.Println("luas :", bangunDatar.luas())
	fmt.Println("keliling :", bangunDatar.keliling())
	fmt.Println("jari-jari :", bangunDatar.(lingkaran).jariJari())
}

// Assert JSON : untuk test ini nama file harus diubah menjadi menggunakan "_test" misalnya file ini menjadi main_test.go
// jalankan langsung dengan => go test

// func TestHitungVolume(t *testing.T) {
// 	assert.Equal(t, 30, 30, "perhitungan volume salah")
// }

// func TestHitungLuas(t *testing.T) {
// 	assert.Equal(t, 15, 15, "Perhitungan luas salah")
// }

// func TestHitungKeliling(t *testing.T) {
// 	assert.Equal(t, 5, 5, "Perhitungan keliling salah")
// }


// Penerapan Interface
func (l lingkaran) jariJari() float64{
	return l.diameter / 2
}

func (l lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.jariJari(), 2)
}

func (l lingkaran) keliling() float64{
	return math.Pi * l.diameter
}

func (p persegi) luas() float64 {
	return math.Pow(p.sisi, 2)
}

func (p persegi) keliling() float64 {
	return p.sisi * 4
}