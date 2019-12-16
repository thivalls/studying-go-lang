package constant

import "fmt"

const (
	b int = 22
	c int = 22
)

func constant() {
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
}