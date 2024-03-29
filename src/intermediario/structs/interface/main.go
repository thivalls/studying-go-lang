package main

import "fmt"

type EntityInterface interface {
	start() string
}

type Names []interface{}

type Entity struct {
	Name string
	Year int
	Color string
}

type Car struct {
	Entity
	Type string
}

type Motorcycle struct {
	Entity
	Type string
}

func (c Car) start() string {
	return "ligando " + c.Name
}

func (m Motorcycle) start() string {
	return "ligando " + m.Name
}

func startEntity(e EntityInterface) string {
	return e.start()
} 

func (n *Names) load() {
	*n = Names{
		"Nair",
		"Paulo",
		"Thiago",
		"ANdre",
	}
}

func (n Names) printNames() {
	fmt.Println(n)
}

func main() {
	// car := Car{
	// 	Entity: Entity {
	// 		Name: "Corolla",
	// 		Year: 2020,
	// 		Color: "verde",
	// 	},
	// 	Type: "Carro",
	// }

	// moto := Motorcycle{
	// 	Entity: Entity {
	// 		Name: "CG150",
	// 		Year: 2020,
	// 		Color: "vermelha",
	// 	},
	// 	Type: "Moto",
	// }

	// fmt.Println(startEntity(car))
	// fmt.Println(startEntity(moto))

	var names Names
	var outros Names

	names.load()
	names.printNames()

	outros.load()
	outros.printNames()
}

