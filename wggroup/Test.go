package wggroup

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var wg1 sync.WaitGroup

func Test() {
	fmt.Printf("%T------wg=%v------wg1=%v", wg == wg1, wg, wg1)
}
