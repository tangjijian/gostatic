package reflect

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

type Abc struct {
}

var abc Abc

func (a Abc) Get(val int) {
	fmt.Println("abc")
}

func Test() {

	//var books = 6
	//var author = "JoinSmith"
	//var isbook = true
	funcType := reflect.TypeOf(abc.Get)
	funcObj := runtime.FuncForPC(reflect.ValueOf(abc.Get).Pointer())
	if funcObj == nil {
		panic("cannot find the method")
	}
	funcNameSli := strings.Split(funcObj.Name(), ".")
	lFuncSli := len(funcNameSli)

	//fmt.Printf("类型=%T,类型名=%s\n", reflect.TypeOf(books), reflect.TypeOf(books))
	//fmt.Printf("类型=%T,类型名=%s\n", reflect.TypeOf(author), reflect.TypeOf(author))
	//fmt.Printf("类型=%T,类型名=%s\n", reflect.TypeOf(isbook), reflect.TypeOf(isbook))
	fmt.Printf("类型=%T,类型名=%s\n", reflect.TypeOf(abc), reflect.TypeOf(abc.Get))
	fmt.Printf("%d,%T\n", lFuncSli, lFuncSli)
	fmt.Println(funcObj.Name())
	fmt.Println(funcNameSli)
	fmt.Printf("funcType=%T,%v\n", funcType.NumIn(), funcType.NumIn())
	fmt.Printf("funcObj====%v,%T\n", funcObj, funcObj)
}

func SetVal(x interface{}) {
	value := reflect.ValueOf(x)
	if value.Kind() == reflect.String {
		value.SetString("反射一下")
	}
}
func SetVal1(x interface{}) {
	value := reflect.ValueOf(x)
	if value.Elem().Kind() == reflect.String {
		value.Elem().SetString("指针反射一下")
	}
}
