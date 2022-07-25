package pointer

import "fmt"

func GetPtr() (int, *int) {

	var a = 10
	var pa = &a

	return a, pa
}

func TestPtr() {

	var p *int

	fmt.Printf("%v\n", p)
	fmt.Printf("%#v\n", p)
}
