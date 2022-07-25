package constant

import "unsafe"

func GetVal() (string, int, int, uintptr) {
	var arr = [5]int{1, 2, 3, 4, 5}
	const (
		a     = "abc"
		b int = len(a)
		c     = cap(arr)
		d     = unsafe.Sizeof(a)
	)
	return a, b, c, d
}
