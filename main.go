package main

import (
	"bytes"
	"fmt"
	"github.com/aofei/cameron"
	"image/png"
	"math/rand"
	"net"
	"net/http"
	"runtime"
	"static/channel"
	"static/constant"
	"static/goroutinePool"
	"static/goroutineQps"
	"static/gotostatment"
	"static/interfacestatment"
	"static/pointer"
	"static/reflect"
	"static/selectstatment"
	"static/struction"
	"static/switchstatment"
	"static/tcpip"
	"static/ticker"
	"static/timer"
	"static/wggroup"
	"time"
	"unsafe"
)

func Hello() {
	defer tcpip.Wg.Done()
	fmt.Println("Hello World")
}
func main() {
	goroutineQps.Init()
}
func main19() {
	selectstatment.CheckIsFull()
}
func main18() {
	selectstatment.RandChan()
}
func main17() {
	//selectstatment.Init()
	selectstatment.TestDeadLock()
	time.Sleep(time.Second * 2)
}
func main16() {
	ticker.Init()

	time.Sleep(time.Second * 5)
}
func main15() {

	timer.Static()

	time.Sleep(time.Second * 2)
}
func main14() {
	goroutinePool.Init()
}
func main13() {
	var ch chan int
	ch1 := make(chan int)
	fmt.Printf("%T,%v", ch, ch)
	fmt.Printf("%T,%v", ch1, ch1)
	channel.Init()
}
func main12() {
	channel.ChanClose()
}
func main11() {
	channel.SyncChan()
	time.Sleep(time.Second * 2)
}
func main10() {
	wggroup.Test()
}
func main9() {
	go func() {
		defer fmt.Println("A.defer")
		func() {
			defer fmt.Println("B.defer")
			// 结束协程
			runtime.Goexit()
			defer fmt.Println("C.defer")
			fmt.Println("B1")
		}()
		fmt.Println("A1")
	}()
	for {
	}
}
func main8() {
	go func(s string) {
		for i := 0; i < 2; i++ {
			fmt.Println(s)
		}
	}("world")
	// 主协程
	for i := 0; i < 2; i++ {
		// 切一下，再次分配任务
		runtime.Gosched()
		fmt.Println("hello")
	}
}
func main7() {
	coon, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		return
	}
	tcpip.Wg.Add(1)
	go tcpip.ProcessClient(coon)
	tcpip.Wg.Wait()
}
func main6() {
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen failed err", err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept failed err:", err)
			continue
		}
		go tcpip.Process(conn)
	}
}
func main5() {
	interfacestatment.Add(interfacestatment.Controller{})

}
func main4() {
	reflect.Test()
}
func main3() {
	tcpip.Wg.Add(1)
	go Hello()
	fmt.Println("go 并发")
	tcpip.Wg.Wait()
}
func main2() {

	//student := struction.Student{}
	//student.Set("JoinSmith")
	//fmt.Println(student.Get())
	//switchstatment.Do1()
	//selectstatment.Do()
	animals := []interfacestatment.Animals{interfacestatment.Dog{}, interfacestatment.Cat{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
	res, err := gotostatment.MakeErr()
	if err != nil {
		goto Forward
	} else {
		goto Exit
	}
Exit:
	fmt.Println("退出")
Forward:
	fmt.Println(res)
	//function.GetArrLen()
	//function.GetAttCap()
	//fmt.Println(function.GetSquareRoot(9))
	//nextNumber := function.GetSequence()
	//fmt.Println(nextNumber())
	//fmt.Println(nextNumber())
	//fmt.Println(nextNumber())
	//fmt.Println(nextNumber())
	//function.TestDefer()
	//arrays.Set()
	//arrays.Get()
	//arr := [2]int{1, 2}
	//arrays.TransParams(&arr)
	//fmt.Println(arr)
	//arr := [2]int{1, 2}
	//arrays.TransParams(arr)
	//fmt.Println(arr)
	//pointer.TestPtr()
	//jinmiao := struction.Jinmao{}
	//jinmiao.House = struction.DogHouse{}
	//jinmiao.House.Wide = 30000
	//jinmiao.House.Height = 40000
	//struction.MakeHouse(&jinmiao.House)
	//jinmiao.ScanHouse()
	//fmt.Printf("狗窝的宽为%d,长为%d\n", jinmiao.House.Wide, jinmiao.House.Height)
	//s := slicestatment.GetSlice()
	//fmt.Println(len(s))
	//fmt.Println(cap(s))
	//
	//for i := 0; i < 15; i++ {
	//	s = append(s, i)
	//}
	//fmt.Println(len(s))
	//fmt.Println(cap(s))
	//for i := 0; i < 15; i++ {
	//	fmt.Printf(" %d", s[i])
	//}
	//fmt.Println(s[1:3])
	//slicestatment.EditSlice()
	//slicestatment.CpSlice()
	//rangestatment.GetVal()
	//mapstatment.TestMap()
	//fmt.Printf("%d ", recursion.Test(15))
	//var phone interfacestatment.Phone
	//
	//phone = new(interfacestatment.Iphone)
	//fmt.Println(phone.Call())
	//phone = new(interfacestatment.NokiaPhone)
	//fmt.Println(phone.Call())
	//reflect.Test()
	address := ""
	//reflect.SetVal(address)
	reflect.SetVal1(&address)
	fmt.Println(address)

}
func identicon(rw http.ResponseWriter, req *http.Request) {
	//http.ListenAndServe("localhost:8080", http.HandlerFunc(identicon))
	buf := bytes.Buffer{}
	png.Encode(&buf, cameron.Identicon([]byte(req.RequestURI), 540, 60))
	rw.Header().Set("Content-Type", "image/png")
	buf.WriteTo(rw)
}
func main1() {
	//variable.SayHello()
	//_, b, c := variable.Variable()
	//
	//for i := 0; i < c; i++ {
	//	fmt.Println(b)
	//}
	fmt.Println(constant.GetVal())
	fmt.Println(2 << 2)
	fmt.Println(2 << 3)
	fmt.Println(2 << 4)
	fmt.Println(2 << 5)

	fmt.Println(64 >> 5)

	type Cat struct {
		voice string
		eat   string
	}
	cat := Cat{voice: "喵"}
	type WhiteCat struct {
		name string
		age  int
		cat  Cat
	}
	white := WhiteCat{name: "Tom", age: 18, cat: cat}

	fmt.Println(white.cat.voice)

	var dog struction.Dog
	dog.Sound = "wang"
	var jinmao = struction.Jinmao{Name: "小黑"}
	fmt.Println(dog)
	fmt.Println(jinmao)

	a, pa := pointer.GetPtr()
	fmt.Println(a, unsafe.Sizeof(pa))
	fmt.Printf("指针类型%T\n", pa)
	fmt.Printf("变量a的类型为%T\n", a)

	switchstatment.Do()

}

// 数据生产者
func producer(header string, channel chan<- string) {
	// 无限循环, 不停地生产数据
	for {
		// 将随机数和字符串格式化为字符串发送给通道
		channel <- fmt.Sprintf("%s: %v", header, rand.Int31())
		// 等待1秒
		time.Sleep(time.Second)
	}
}

// 数据消费者
func customer(channel <-chan string) {
	// 不停地获取数据
	for {
		// 从通道中取出数据, 此处会阻塞直到信道中返回数据
		message := <-channel
		// 打印数据
		fmt.Println(message)
	}
}

var var1, var2 = 1, 2

//http.proxy=http://tangjijian:123456@git.petope.com:80
