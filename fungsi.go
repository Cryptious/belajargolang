package main

import "fmt"

func main() {
	// fmt.Println(tampilkan_pesan(10, 5))

	var x1, x2, x3 = (tampilkan_pesan(10, 5))
	fmt.Println(x1)
	fmt.Println(x2)
	fmt.Println(x3)

}

func tampilkan_pesan(x int, y int) (int, int, int) {
	var z = x * y
	var a = x / y
	var b = x + y
	return z, a, b
}
