package goroutineQps

import (
	"fmt"
	"github.com/petermattis/goid"
	"sync"
	"time"
)

var x int64
var wg sync.WaitGroup
var lock sync.Mutex
var rwlock sync.RWMutex

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

// 读写锁
func write() {

	rwlock.Lock()
	x = x + 1
	time.Sleep(time.Millisecond * 10)
	rwlock.Unlock()
	wg.Done()
}
func read() {
	rwlock.Lock()
	time.Sleep(time.Millisecond)
	rwlock.Unlock()
	wg.Done()
}

func RWInit() {

	start := time.Now()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}
	end := time.Now()
	wg.Wait()
	fmt.Println(end.Sub(start))
}
