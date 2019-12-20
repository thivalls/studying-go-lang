package semaphores

import "fmt"

func semaphores() {
	channel := make(chan int)
	next := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- i
		}
		next <- true
	}()

	go func() {
		for i := 0; i < 20; i++ {
			channel <- i
		}
		next <- true
	}()

	go func() {
		<-next
		<-next
		close(channel)
		close(next)
	}()

	for number := range(channel) {
		fmt.Println(number)
	}
}