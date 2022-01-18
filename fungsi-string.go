package main

import (
	"fmt"
	"strings"
)

func main() {
	// var apakahada = strings.Contains("disini adalah rumah", "ada") // yang mengandung
	// var apakahada = strings.HasPrefix("disini adalah rumah", "disini") // awalan
	// var apakahada = strings.HasSuffix("disini adalah rumah", "lah") // akhiran
	// fmt.Println(apakahada)

	// var berapabanyak = strings.Count("misalnya ini tulisan", "i") // menghitung banyaknya char
	// fmt.Println(berapabanyak)

	// var index1 = strings.Index("misalnya ini tulisan", "l") // mengetahui posisi char
	// fmt.Println(index1)

	var text = "buah apel"
	var cari = "a"
	var text_baru = strings.Replace(text, cari, "u", 1) // (text_yg_ingin_diubah,str_pengganti,str_yg_ingin_diganti,brp_banyak)
	fmt.Println(text_baru)
}
