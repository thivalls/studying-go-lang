package slices

import "fmt"

func slices() {
	// slice := make([]int, 5, 8)
	// clone := slice
	// clone = append(clone, 5, 6, 7, 9)
	// slice[0] = 555
	slice := []string {
		"Thiago",
		"Thiago",
		"valls",
		"Bertola",
	}

	fmt.Println(slice[1:3])
	// fmt.Println(clone)
}