package main

import (
	"fmt"
	"net/url"
)

func main() {
	var url_text = "http://google.com:8080/hai/?cari=buku favorit&author=Martino&tahun=2010"
	var u, err = url.Parse(url_text)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("url :", url_text)
	fmt.Println("host nya : ", u.Host)
	fmt.Println("Path :", u.Path)

	var cari = u.Query()["cari"][0]
	var author = u.Query()["author"][0]
	var tahun = u.Query()["tahun"][0]

	fmt.Println("Cari :", cari)
	fmt.Println("Author :", author)
	fmt.Println("Tahun :", tahun)
}
