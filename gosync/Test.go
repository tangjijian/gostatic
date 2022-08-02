package gosync

import (
	"fmt"
	"image"
	"strconv"
	"sync"
)

var icon map[string]image.Image
var loadOnce sync.Once

func Init(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("init function")
}

func loadIcons() {

	icon = map[string]image.Image{}
}

// load 是并发安全的
func load(name string) image.Image {
	if icon == nil {
		loadOnce.Do(loadIcons)
	}
	return icon[name]
}

// Go语言中内置的map不是并发安全的
var m map[string]int

var wg sync.WaitGroup

func set(k string, v int) {
	m[k] = v
}
func get(k string) int {
	return m[k]
}
func ThreadUnSafeMap() {
	//m = make(map[string]int) // 不是并发安全的

	//for i := 0; i < 200; i++ {
	//	wg.Add(1)
	//	go func(n int) {
	//		key := strconv.Itoa(i)
	//		set(key, i)
	//		fmt.Printf("key=%v,v=%v\n", key, get(key))
	//		wg.Done()
	//	}(i)
	//}
	m := sync.Map{}
	for i := 0; i < 200; i++ {
		wg.Add(1)
		k := strconv.Itoa(i)
		m.Store(k, i)
		value, _ := m.Load(k)
		fmt.Printf("key=%v,v=%v\n", k, value)
		wg.Done()
	}
	wg.Wait()

}
