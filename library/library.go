package library

import "fmt"

func iniprivate() {
	fmt.Println("saya di Private")
}

func Inipublic() {
	fmt.Println("Saya di Public")
	iniprivate()
}
