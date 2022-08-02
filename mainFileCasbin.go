package main

import (
	"fmt"
)

func test(params ...interface{}) {
	fmt.Println(params)
}
func main() {
	test(1, 2, 3)

}
