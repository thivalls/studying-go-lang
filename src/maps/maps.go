package maps

import "fmt"

func maps() {
	m := make(map[string]int)

	m["a"] = 10
	m["b"] = 20
	m["c"] = 30

	delete(m, "a")

	// value, exists := m["a"]

	// if exists { 
	// 	fmt.Println(value)
	// } else {
	// 	fmt.Println(value)
	// }

	// var x = map[string]int{"x":10, "y": 11}
	// x1 := map[string]int{"x":10, "y": 11}

	// fmt.Println(x, x1)

	if value, exists := m["a"]; exists {
		fmt.Println(value)
	} else {
		fmt.Println(exists, " ops !")
	}
}