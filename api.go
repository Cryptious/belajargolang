package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "mysql-master"
	"net/http"
)

type menu_makanan struct {
	Id        string
	Nama_menu string
	Harga     int
}

func koneksi() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(localhost)/db_menumakanan")

	if err != nil {
		return nil, err
	}
	return db, nil
}

var data []menu_makanan

/*
var data = []menu_makanan{
	menu_makanan{"C01", "Karedok", 20000},
	menu_makanan{"C02", "Pecel Lele", 15000},
	menu_makanan{"C03", "Ketoprak", 17000},
	menu_makanan{"C04", "Ayam Bakar", 25000},
}
*/

func main() {
	ambil_data()
	http.HandleFunc("/menu", ambil_menu)
	http.HandleFunc("/cari_menu", cari_menu)
	fmt.Println("menjalankan Web server")
	http.ListenAndServe(":8080", nil)
}

func ambil_menu(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		var result, err = json.Marshal(data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}

func cari_menu(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")

	if r.Method == "POST" {
		var namamenu = r.FormValue("Nama_menu")
		var result []byte
		var err error

		for _, each := range data {
			if each.Nama_menu == namamenu {
				result, err = json.Marshal(each)

				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				w.Write(result)
				return
			}
		}
		http.Error(w, "Menu Makanan Tidak Tersedia", http.StatusBadRequest)
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}

func ambil_data() {
	db, err := koneksi()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query("Select * from tb_menu")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		var each = menu_makanan{}
		var err = rows.Scan(&each.Id, &each.Nama_menu, &each.Harga)

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		data = append(data, each)
	}
	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

}
