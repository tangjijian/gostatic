package slicestatment

import (
	"fmt"
)

func GetSlice() []int {
	s := make([]int, 0, 10)
	return s
}

func EditSlice() {
	arr := [5]int{1, 2, 3, 4, 5}
	s := arr[:]
	s = append(s, 6)
	fmt.Printf("原始数组=%v\n", arr)
	fmt.Printf("切片=%v", s)
}
func CpSlice() {

	var s []int
	PrintSlice(s)
	s = append(s, 1)
	PrintSlice(s)
	s = append(s, 2, 3, 4, 5)
	PrintSlice(s)
	s1 := make([]int, len(s), cap(s)*2)
	copy(s1, s)
	PrintSlice(s1)
}
func PrintSlice(s []int) {
	fmt.Printf("len = %d,cap=%d,val=%v\n", len(s), cap(s), s)
}
