package deadlock

import "fmt"

func deadlock() {

	channel := make(chan int)
	
	go func(){
		channel <- 10
	}()

	fmt.Println(<-channel)

}