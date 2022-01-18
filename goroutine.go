package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	go tampilpesan(5, "saya kirim")
	tampilpesan(5, "saya kedua")

	var input string
	fmt.Scanln(&input)
}

func tampilpesan(x int, pesan string) {
	for i := 0; i < x; i++ {
		fmt.Println((i + 1), pesan)
	}
}
