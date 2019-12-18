package forExample

import "fmt"

func forExample() {
	// FOR
	// for i := 0; i <= 5; i++ {
	// 	fmt.Println(i)
	// }

	// WHILE
	// w := 0
	// for w < 5 {
	// 	fmt.Println(w)
	// 	w++
	// }

	end := 100
	x := 0

	for {
		fmt.Println(x)
		
		if x == end {
			break
		}

		x++
	}
}