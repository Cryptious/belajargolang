package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

var baseURL = "http://localhost:8080"

type menu_makanan struct {
	Id        string
	Nama_menu string
	Harga     int
}

func ambil_api(menu string) (menu_makanan, error) {
	var err error
	var client = &http.Client{}
	var data menu_makanan

	var param = url.Values{}
	param.Set("Nama_menu", menu)
	var payload = bytes.NewBufferString(param.Encode())

	request, err := http.NewRequest("POST", baseURL+"/cari_menu", payload)

	if err != nil {
		return data, err
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response, err := client.Do(request)

	if err != nil {
		return data, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}

func main() {
	var menu, err = ambil_api("Bakso Malang")
	if err != nil {
		fmt.Println("Menu tidak tersedia", err.Error())
		return
	}
	fmt.Println("ID: ", menu.Id, "Menu: ", menu.Nama_menu, "Harga: ", menu.Harga)

}
