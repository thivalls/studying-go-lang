package main

import "fmt"

type Car struct {
	Name string
	Year int
	Color string
}

func (c Car) carInfo() string {
	return fmt.Sprintf("Nome: %s\n Ano: %d\n Cor: %s", c.Name, c.Year, c.Color)
}

func main() {
	car1 := Car{Name: "Corolla", Year: 2020, Color: "Red"}

	fmt.Println(car1.carInfo())
}