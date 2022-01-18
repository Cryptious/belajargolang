package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	var text = "text ini rahasia"
	var sha = sha1.New()
	sha.Write([]byte(text))
	var enkripsi = sha.Sum(nil)
	var stringenkripsi = fmt.Sprintf("%x", enkripsi)

	fmt.Println(stringenkripsi)
}
