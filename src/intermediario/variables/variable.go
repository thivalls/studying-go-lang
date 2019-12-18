package variable

import "fmt"

var b int = 22

func variable() {
	a := 10
	b := 10.45
	c := "Hello"
	d := false
	e := ` escrever mais linha `
	f := 't'

	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)
	fmt.Printf("%T \n", e)
	fmt.Printf("%T \n", f)
}