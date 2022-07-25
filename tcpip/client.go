package tcpip

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

var Wg sync.WaitGroup

func ProcessClient(conn net.Conn) {
	defer conn.Close()
	defer Wg.Done()
	inputReader := bufio.NewReader(os.Stdin)

	for {

		input, _, _ := inputReader.ReadLine()

		inputInfo := string(input)

		if strings.ToUpper(inputInfo) == "Q" {
			return
		}

		_, err := conn.Write([]byte(inputInfo))

		if err != nil {
			fmt.Println(err)
			return
		}
		buf := [512]byte{}
		fmt.Println("a")
		n, err := conn.Read(buf[:]) // 空数据不能读取，会一直阻塞
		fmt.Println("b")
		if err != nil {
			fmt.Println("Receive failed err : ", err)
			return
		}

		fmt.Println("服务端发来数据:", string(buf[:n]))

	}
}
