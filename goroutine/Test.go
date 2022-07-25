package goroutine

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func Hello() {
	defer wg.Done()
	fmt.Println("Hello World")
}
