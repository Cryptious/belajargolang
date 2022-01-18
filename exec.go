package main

import (
	"fmt"
	"os/exec"
)

func main() {

	var hasil, _ = exec.Command("help").Output()
	fmt.Println(string(hasil))
}
