package ifs

import "fmt"

func ifs() {
	a := 3

	if a > 10 {
		fmt.Println("A > 10")
	} else if a > 5 {
		fmt.Println("A > 5")
	} else {
		fmt.Println("nenhuma opção")
	}

	b := true

	if x := "Cool"; b {
		fmt.Println(x)
	}
}