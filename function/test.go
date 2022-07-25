package function

import (
	"fmt"
	"math"
)

var arr = [5]int{1, 2, 3, 4}

func GetArrLen() {

	fmt.Println(len(arr))
}

func GetAttCap() {
	fmt.Println(cap(arr))
}

var GetSquareRoot = func(x float64) float64 {
	fmt.Printf("%T\n", x)
	return math.Sqrt(x)
}

func GetSequence() func() int {

	i := 0
	return func() int {
		i += 1
		return i
	}
}

// TestDefer defer 语句先进后出
func TestDefer() {
	fmt.Println("开始了")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("结束了")
}
