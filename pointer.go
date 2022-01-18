package main

import "fmt"

func main() {
	var nomor int = 10
	var tas = "Doraemon"
	var alamat_tas *string = &tas
	var alamat_nomor *int = &nomor

	fmt.Println("Nilai dari nomor : ", nomor)
	fmt.Println("Alamat Variabel Nomor :", alamat_nomor)
	fmt.Println("Alamat Variabel tas :", alamat_tas)
}
