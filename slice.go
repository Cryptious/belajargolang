package main

import "fmt"

func main() {
	var buah = []string{"Apel", "Mangga", "Pisang", "Jambu"}
	buah = append(buah, "Durian")
	fmt.Println(buah)
	fmt.Println("Jumlah Element : ", len(buah))

}
