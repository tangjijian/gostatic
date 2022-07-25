package channel

import "fmt"

func receive(ch chan int) {
	rev := <-ch
	fmt.Println("接收到的值:", rev)
}
func SyncChan() {
	ch := make(chan int, 1)
	go receive(ch)
	ch <- 10
	fmt.Println("发送成功")
}

func ChanClose() {

	ch := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}()
	for {
		if data, ok := <-ch; ok {
			fmt.Println(data)
		} else {
			break
		}
	}
	fmt.Println("main 结束")
}

func counter(out chan<- int) {
	for i := 0; i < 20; i++ {
		out <- i
	}

	close(out)
}
func square(out chan<- int, in <-chan int) {
	for data := range in {
		out <- data
	}
	close(out)
}
func printer(out <-chan int) {
	for data := range out {
		fmt.Println(data)
	}
}
func Init() {
	ch1 := make(chan int, 20)
	ch2 := make(chan int, 20)

	go counter(ch1)
	go square(ch2, ch1)
	printer(ch2)
}
