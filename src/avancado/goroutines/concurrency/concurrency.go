package concurrency

import (
	"fmt"
	"time"
	"math/rand"
	"sync"
	"runtime"
)

var waitGroup sync.WaitGroup

var result int

func init() {
	// versoes mais novas utilizam o máximo de cpus, porém podemos setar manualmente
	// se tiver apenas um CPU roda de forma concorrente, mas se tiver mais roda paralelamente
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func concurrency() {
	waitGroup.Add(2)
	go runProcess("P1", 20)
	go runProcess("P2", 20)
	waitGroup.Wait()
}

func runProcess(name string, total int) {
	for i := 0; i < total; i++ {
		t := time.Duration(rand.Intn(255))
		time.Sleep(time.Millisecond * t)
		fmt.Println(name, " -> ", i, "partial result is: ", result)
	}
	waitGroup.Done()
}