package main

import (
	"fmt"
	"strconv"
)

func main() {
	defer pulihkan_saya()
	var input string
	fmt.Print("Masukan Angka :")
	fmt.Scanln(&input)

	var number int
	var err error
	number, err = strconv.Atoi(input)

	if err == nil {
		fmt.Println(number, "ini adalah angka")
	} else {
		fmt.Println(input, "Ini Bukan angka")
		// fmt.Println(err.Error())
		panic(err.Error())
		fmt.Println("Munculkan Saya")
	}
}

func pulihkan_saya() {
	if r := recover(); r != nil {
		fmt.Println("Akhirnya saya ditampilkan")
	} else {
		fmt.Println("Ini Lancar Saja")
	}
}
