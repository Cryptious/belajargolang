package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Nama  string `jason:"Nama"`
	Umur  int
	Kelas string `json:"Kelas"`
}

func main() {
	/* decode data dari Json ke object
	var jsonString = `{"Nama": "Ilham", "Umur": 18, "Kelas": "XII"}`
	var jsonData = []byte(jsonString)

	var data User
	var err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Nama :", data.Nama)
	fmt.Println("Umur :", data.Umur)
	fmt.Println("Kelas :", data.Kelas)
	*/

	// dari object ke  json
	var object = []User{{"Ilham", 18, "XII"}, {"Abdul", 21, "XI"}}
	var jsonData, err = json.Marshal(object)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var jsonString = string(jsonData)
	fmt.Println(jsonString)
}
