package pointers

import "fmt"

func xpto(a *string) string {
	fmt.Println(a)
	fmt.Println(*a)

	*a = *a + " thiago"
	return *a
}

func pointers() {
	c := "teste"

	// fmt.Println(&c)

	// d := &c

	// *d = 30

	// fmt.Println(*d)
	// fmt.Println(c)

	// var x *int = &c

	//fmt.Println(c)

	fmt.Println(xpto(&c))

	//fmt.Println(c)
}