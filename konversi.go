package main

import (
	"fmt"
	"strconv"
)

func main() {

	//STR to INT
	// var str = "156"
	// var num, err = strconv.Atoi(str) // Str to Int
	// if err == nil {
	// 	fmt.Println(num)
	// 	fmt.Println(num + 10) // test String or Int
	// }

	// INT to STR

	// var num = 156
	// var str = strconv.Itoa(num) //Int to STR
	// fmt.Println(str)

	// str to float
	// var num = "30.7"
	// var flt, err = strconv.ParseFloat(num, 16)
	//
	// if err == nil {
	// 	fmt.Println(flt)
	// 	fmt.Println(flt + 10)
	// }

	// Float to Str
	var flt = 30.7
	var str = strconv.FormatFloat(flt, 'f', 6, 64)

	fmt.Println(str)
	// fmt.Println(str + 10) // pengujian

}
