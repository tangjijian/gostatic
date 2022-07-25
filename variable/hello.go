package variable

import "fmt"

func SayHello() {
	fmt.Println("hello world")
}
func Variable() (string, string, int) {

	var a = "W3CScholl"
	var b string = a + "是个好地方"
	c := 10
	return a, b, c
}
