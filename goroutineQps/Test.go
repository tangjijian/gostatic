package goroutineQps

import (
	"fmt"
	"github.com/petermattis/goid"
	"sync"
)

var x int64
var wg sync.WaitGroup
var lock sync.Mutex

func add() {
	for i := 0; i < 10000; i++ {
		x = x + 1
	}
	wg.Done()
}
func Init() {
	wg.Add(2)
	//go add()
	//go add()
	go MutexAdd()
	go MutexAdd()
	wg.Wait()
	fmt.Println(x)
}
func MutexAdd() {
	for i := 0; i < 10000; i++ {
		lock.Lock()
		fmt.Println(goid.Get())
		x = x + 1
		lock.Unlock()
	}
	wg.Done()
}
