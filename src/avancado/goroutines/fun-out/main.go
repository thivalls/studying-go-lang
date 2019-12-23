package main

import "fmt"

func main() {

	go func() {
		c := generate(5, 10)
		
		d1 := <-divide(c)
		d2 := <-divide(c)

		fmt.Println("Finalized...", d1, d2)
	}()

}

func generate(numbers ...int) chan int {
	channel := make(chan int)
	for _, n := range(numbers) {
		channel <- n
	}
	return channel
}

func divide(input chan int) chan int {
	channel := make(chan int)
	
	go func() {
		for number := range(input) {
			channel <- number / 2
		}
	}()

	return channel
}
