package functionsExample

import "fmt"

func returnNameLastName(name string, lastname string) (string, string) {
	return name, lastname
}

func functionsExample() {
	name, last := returnNameLastName("Thiago", "Valls");

	fmt.Printf("%s %s \n", name, last + " é o nome indicado pela função")
}