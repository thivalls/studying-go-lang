package main

import (
	"fmt"
	"time"
	"math/rand"
	"sync"
	"runtime"
)

var waitGroup sync.WaitGroup

func init() {
	// versoes mais novas utilizam o máximo de cpus, porém podemos setar manualmente
	// se tiver apenas um CPU roda de forma concorrente, mas se tiver mais roda paralelamente
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	waitGroup.Add(2)
	go runProcess("P1", 10)
	go runProcess("P2", 10)
	waitGroup.Wait()
}

func runProcess(name string, total int) {
	for i := 0; i < total; i++ {
		fmt.Println(name, " -> ", i)
		t := time.Duration(rand.Intn(255))
		time.Sleep(time.Millisecond * t)
	}
	waitGroup.Done()
}