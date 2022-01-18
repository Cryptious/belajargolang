package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var nama_saya = "Marutino"

	var encodeString = base64.StdEncoding.EncodeToString([]byte(nama_saya))
	fmt.Println("Encoding String :", encodeString)

	var decodeByte, _ = base64.StdEncoding.DecodeString(encodeString)
	var decodeString = string(decodeByte)
	fmt.Println("Sesudah decode : ", decodeString)
}
