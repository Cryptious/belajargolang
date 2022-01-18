package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("Hai")
	fmt.Println("Selamat Datang")
	fmt.Println("Di Rumah Saya")
	os.Exit(1)
	tampilkan()
}

func tampilkan() {
	fmt.Println("Saya ditampilkan")
}
