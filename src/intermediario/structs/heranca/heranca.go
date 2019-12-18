package heranca

import "fmt"

type Car struct {
	Name string
	Year int
	Color string
}

type SuperCar struct {
	Car
	CanFly bool
	Name string
}

func (c Car) carInfo() string {
	return fmt.Sprintf("Nome: %s\n Ano: %d\n Cor: %s", c.Name, c.Year, c.Color)
}

func (s SuperCar) carInfo() string {
	return fmt.Sprintf("Informações override - Nome: %s\n Ano: %d\n Cor: %s", s.Name, s.Year, s.Color)
}

func heranca() {
	car1 := SuperCar{
		Car: Car{Name: "Corolla", Year: 2020, Color: "Red"},
		CanFly: true,
		Name: "Uno",
	}

	fmt.Println(car1.carInfo())
}