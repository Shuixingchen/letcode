package programMod

import (
	"fmt"
	"time"
)

//扇入, 多个生产者，一个消费者

//生产者，把数据写入inch
func producer(inch chan<- int) {
	var i int
	for {
		inch <- i
		time.Sleep(1 * time.Second)
		i++
	}
}

//消费者从out读取数据
func consumer(out <-chan int) {
	for item := range out {
		fmt.Println(item)
	}
}

func FanIn() {
	inch := make(chan int, 0)
	outch := make(chan int, 0)
	go producer(inch)
	go producer(inch)
	go consumer(outch)
	for i := range inch {
		outch <- i
	}
}
