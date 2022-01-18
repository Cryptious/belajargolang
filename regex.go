package main

import (
	"fmt"
	"regexp"
)

func main() {
	var text = "jeruk1 pisang9 anggur8"
	var regex, err = regexp.Compile(`[a-z]+\d`)

	if err != nil {
		fmt.Println(err.Error())
	}
	var hasil = regex.FindAllString(text, -1)
	fmt.Println(hasil)

	var cocok = regex.MatchString(text)
	fmt.Println(cocok)

	var index1 = regex.FindStringIndex(text)
	fmt.Println(index1)

	var new_string = regex.ReplaceAllString(text, "durian")
	fmt.Println(new_string)
}
