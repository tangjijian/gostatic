package arrays

import "fmt"

var arr [4][5]int

func Set() [4][5]int {
	for i := 0; i < 4; i++ {
		for j := 0; j < 5; j++ {
			arr[i][j] = i + j
		}
	}
	return arr
}
func Get() {
	for i := 0; i < 4; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("%d", arr[i][j])
		}
		fmt.Printf("\n")
	}
}

func TransParamsRefer(arr *[2]int) {
	arr[1] = 5
	fmt.Printf("%T\n", arr)
}

func TransParams(arr [2]int) {
	arr[1] = 5
	fmt.Printf("%T\n", arr)
}
