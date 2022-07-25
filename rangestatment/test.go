package rangestatment

import "fmt"

func GetVal() {
	str := "你好世界"

	for i, v := range str {
		fmt.Printf("%d,%c\n", i, v)
	}
}
