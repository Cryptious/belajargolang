package main

import (
	"database/sql"
	"fmt"
	_ "mysql-master"
)

type pelajar struct {
	id    int
	nama  string
	umur  int
	kelas int
}

func koneksi() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(localhost)/belajar_golang")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func sql_tampil() {
	db, err := koneksi()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
	// var kelas = 2
	rows, err := db.Query("Select * from tb_pelajar")

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []pelajar

	for rows.Next() {
		var each = pelajar{}

		var err = rows.Scan(&each.id, &each.nama, &each.umur, &each.kelas)

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		result = append(result, each)
	}
	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, each := range result {
		fmt.Println(each.id, each.nama, each.umur, each.kelas)
	}

}

/* menambah data
func sql_tambah() {
	db, err := koneksi()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
	_, err = db.Exec("insert into tb_pelajar values (?,?,?,?)", nil, "Irene", 19, 2)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("tambah data berhasil")
}
*/

/* mengubah data
func ubah_data() {
	db, err := koneksi()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
	_, err = db.Exec("update tb_pelajar set umur = ? where nama = ?", 20, "Fenny")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Ubah Berhasil")
}
*/

/* hapus data
func hapus_data() {
	db, err := koneksi()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
	_, err = db.Exec("delete from tb_pelajar where id =? ", 5)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("hapus Berhasil")
}
*/

func main() {
	// sql_tambah()
	// ubah_data()
	// hapus_data()
	sql_tampil()
}
