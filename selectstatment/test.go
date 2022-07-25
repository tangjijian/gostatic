package selectstatment

import (
	"fmt"
	"time"
)

func Do() {
	var c1, c2, c3 chan int
	var i1, i2 int
	/*	var a1 chan<- int
		var a2 <-chan int*/
	select {
	case i1 = <-c1:
		fmt.Printf("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Printf("sent ", i2, " to c2\n")
	case i3, ok := (<-c3): // same as: i3, ok := <-c3
		if ok {
			fmt.Printf("received ", i3, " from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}
}

func test1(ch chan string) {
	time.Sleep(time.Second * 5)
	ch <- "test1"
}
func test2(ch chan string) {
	time.Sleep(time.Second * 2)
	ch <- "test2"
}
func Init() {
	output1 := make(chan string)
	output2 := make(chan string)
	go test1(output1)
	go test2(output2)
	select {
	case str1 := <-output1: // 主goroutine中接收nil管道的数据会产生死锁
		fmt.Println("str1==", str1)
	case str2 := <-output2:
		fmt.Println("str2==", str2)
	}
}
func causeDead(ch chan int) {
	<-ch
}
func TestDeadLock() {
	var ch1 chan int
	fmt.Printf("%v\n", ch1)
	ch := make(chan int)
	fmt.Printf("%T%v", ch, ch)
	go causeDead(ch)
}
func RandChan() {
	ch1 := make(chan int)
	ch2 := make(chan string)

	go func(ch chan int) {
		ch <- 123
	}(ch1)

	go func(ch chan string) {
		ch <- "abc"
	}(ch2)

	select {
	case num := <-ch1:
		fmt.Println("int=", num)
	case str := <-ch2:
		fmt.Println("string=", str)
	}
}
func CheckIsFull() {

	ch := make(chan string, 10)

	go write(ch)
	for data := range ch {
		fmt.Println(data)
		time.Sleep(time.Second)
	}

}
func write(input chan string) {
	for {

		select {
		case input <- "hello":
			fmt.Println("write hello")
		default:
			fmt.Println("channel was full")
		}
		time.Sleep(time.Millisecond * 500)
	}
}
