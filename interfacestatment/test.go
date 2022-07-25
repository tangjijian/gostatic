package interfacestatment

import "fmt"

type Animals interface {
	Speak() string
}
type ControllerInterface interface {
	Get()
}
type Controller struct {
	appName string
}

func (c Controller) Get() {

}

type Route interface {
	Add(c ControllerInterface)
}
type MyRoute struct {
}

func Add(c ControllerInterface) {
	fmt.Println("接口类型测试")
}

type Phone interface {
	Call() string
}
type Dog struct {
}
type Iphone struct {
}
type NokiaPhone struct {
}

func (dog Dog) Speak() string {
	return "woof!"
}

type Cat struct {
}

func (cat Cat) Speak() string {
	return "maio!"
}
func (iphone Iphone) Call() string {
	return "I am iPhone. I am calling you"
}
func (nokia NokiaPhone) Call() string {
	return "I am nokiaPhone.I am calling you too!"
}

//func NewObj() others.DbObj {
//	return
//}
//func init() {
//	NewObj()
//}
