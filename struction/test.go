package struction

import "fmt"

type Dog struct {
	Sound string
}

//type DogHouse struct {
//	Wide   int
//	Height int
//}
//type Jinmao struct {
//	Name  string
//	House *DogHouse
//}
//
//var dog Jinmao
//
//func (j Jinmao) ScanHouse() {
//	fmt.Printf("内部狗窝的宽为%d,长为%d\n", j.House.Wide, j.House.Height)
//}
//func MakeHouse(house *DogHouse) {
//	house.Wide = 400
//	house.Height = 300
//}

type DogHouse struct {
	Wide   int
	Height int
}
type Jinmao struct {
	Name  string
	House DogHouse
}

var dog Jinmao

func (j Jinmao) ScanHouse() {
	fmt.Printf("内部狗窝的宽为%d,长为%d\n", j.House.Wide, j.House.Height)
}
func MakeHouse(house *DogHouse) {
	house.Wide = 400
	house.Height = 300
}
