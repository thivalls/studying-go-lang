package serialize

import (
	"fmt"
	"encoding/json"
)

type Car struct {
	Name string `json:"Nome"`
	Year int `json:"Ano"`
	Color string `json:"Cor"` 
}

func (c Car) getInfoJsonCar() {
	result, _ := json.Marshal(c)
	fmt.Println(string(result))
}

func getJsonCar() string {
	return `{"Nome":"Corolla","Ano": 2020, "Cor":"Green"}`;
}

func serialize() {
	// car := Car{"Corolla",2020,"Green"}

	// car.getInfoJsonCar();

	var car Car
	data := []byte(getJsonCar())

	json.Unmarshal(data, &car)

	fmt.Println(car)
}