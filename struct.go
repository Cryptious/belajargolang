package main

import "fmt"

func main() {
	var x1 pelajar
	x1.nama = "Hojan"
	x1.kelas = 12
	x1.umur = 17

	fmt.Println("Nama :", x1.nama)
	fmt.Println("Kelas :", x1.kelas)
	fmt.Println("Umur :", x1.umur)
	var x2 pelajar
	x2.nama = "Vilan"
	x2.kelas = 11
	x2.umur = 18
	fmt.Println("Nama :", x2.nama)
	fmt.Println("Kelas :", x2.kelas)
	fmt.Println("Umur :", x2.umur)

}

type pelajar struct {
	nama  string
	kelas int
	umur  int
}
