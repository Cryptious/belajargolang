package main

import "fmt"

func main() {
	var buah = [5]string{"apel", "jeruk", "anggur", "melon"}
	buah[4] = "Mangga"
	fmt.Println("Jumlah Element : ", len(buah))
	fmt.Println("Isi Element : ", buah)
}
