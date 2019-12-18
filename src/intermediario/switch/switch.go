package switchExample

import "fmt"

func switchExample() {
	x := "thiso"

	switch x {
		case "thiso":
			fmt.Println("eeee thiagao hein errou o nome")
		case "epa" :
			fmt.Println("achei")
		default:
			fmt.Println("nenhum encontrado")
	}


}