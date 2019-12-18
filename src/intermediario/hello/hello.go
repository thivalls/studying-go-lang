package hello

import (
	"fmt"
	"github.com/satori/go.uuid"
)

func Hello() {
	a, _ := uuid.NewV4()
	fmt.Printf("Hello World!!! from another package with id = %s \n", a)
}
