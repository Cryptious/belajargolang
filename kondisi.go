package main

import "fmt"

func main() {
	var nilai = 5

	switch nilai {
	case 10:
		fmt.Println("Sempurna")
	case 9:
		fmt.Println("Bagus")
	case 8:
		fmt.Println("Lumayan")
	default:
		fmt.Println("Belum Lulus")
	}

	/* kondisi if statement
	  if nilai == 10 {
			fmt.Println("Lulus dengan sempurna")
		} else if nilai > 7 {
			fmt.Println("Lulus")
		} else if nilai == 6 {
			fmt.Println("Sedikit Lagi")
		} else {
			fmt.Println("Belajar Lagi")
		}
	*/
}
