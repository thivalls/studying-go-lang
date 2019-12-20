package channels

import (
	"fmt"
	"sync"
)

var waitGroup sync.WaitGroup

func channels() {
	// msg := make(chan string)

	// go func() {
	// 	msg <- "Hello World"
	// }()

	// result := <- msg

	// go func() {
	// 	msg <- "Hello World 2"
	// }()

	// result1 := <- msg

	// fmt.Println(result, result1)

	channel := make(chan int)

	

	waitGroup.Add(2)

	go func () {
		for i := 0; i < 10; i++ {
			channel <- i
		}
		waitGroup.Done()
	}()

	go func () {
		for i := 0; i < 10; i++ {
			channel <- i
		}
		waitGroup.Done()
	}()

	go func () {
		waitGroup.Wait()
		close(channel)
	}()

	for number := range channel {
		fmt.Println(number)
	}
}