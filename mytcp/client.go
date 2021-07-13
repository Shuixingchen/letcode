package mytcp

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
)

func ClientRun() {
	var wg sync.WaitGroup
	for i:=0; i<10; i++ {
		go func() {
			wg.Add(1)
			conn, err := net.Dial("tcp", "127.0.0.1:8080")
			if err != nil {
				fmt.Println("客户端建立连接失败")
				return
			}
			cConnHandler(conn, wg)
		}()

	}
	wg.Wait()
}

func cConnHandler(c net.Conn, wg sync.WaitGroup) {
	reader := bufio.NewReader(os.Stdin)
	buf := make([]byte, 1024)
	fmt.Println("请输入客户端请求数据...")
	for {
		input, _ := reader.ReadString('\n')
		input = string(input)
		c.Write([]byte(input))
		cnt, err := c.Read(buf)
		if err != nil {
			fmt.Printf("客户端读取数据失败 %s\n", err)
			continue
		}
		fmt.Print("服务器端回复" + string(buf[0:cnt]))
	}
	wg.Done()
}

