package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"sync"
	"time"
)

type Handler struct {
	TaskCh      chan int
	ResultCh    chan string
	Start       int
	End         int
	wg          sync.WaitGroup
	WorkerNum   int
	InterruptCh chan os.Signal //接受os信号
}

func main() {
	// handler := NewHandler()
	// handler.Init()
	// handler.Serve()
	Close()
}

func NewHandler() *Handler {
	return &Handler{
		TaskCh:      make(chan int),
		ResultCh:    make(chan string, 0),
		Start:       0,
		End:         10,
		WorkerNum:   5,
		wg:          sync.WaitGroup{},
		InterruptCh: make(chan os.Signal),
	}
}
func (h *Handler) Init() {
	runtime.GOMAXPROCS(256)
	if h.Start > h.End {
		panic("start > End")
	}
	signal.Notify(h.InterruptCh, os.Interrupt) //接受系统中断信号，会写入通道
}

func (h *Handler) Serve() {
	//把任务写入队列
	go func(s, e int) {
		for i := s; i <= e; i++ {
			h.TaskCh <- i
		}
		close(h.TaskCh)
	}(h.Start, h.End)

	//创建worker,每个worker循环监听任务队列
	h.wg.Add(h.WorkerNum)
	for j := 0; j < h.WorkerNum; j++ {
		go h.StartWork()
	}

	//处理结果，循环读取结果队列
	go func() {
		for res := range h.ResultCh {
			time.Sleep(1e9)
			fmt.Println("res:", res)
		}
	}()
	//wait是等待worker协程，worker完成后，结果通道也就没有数据写了，要及时关掉
	h.wg.Wait()
	close(h.ResultCh)

}

//worker是死循环，任务读完或者有中断信号要退出
func (h *Handler) StartWork() {
	defer h.wg.Done()
	for {
		select {
		case t, ok := <-h.TaskCh:
			if ok == false {
				return
			}
			res := GetData(t)
			fmt.Println("input res:", res)
			h.ResultCh <- res
		default:
		}
	}
}

func GetData(number int) string {
	return strconv.Itoa(number)
}

func Close() {
	var wg sync.WaitGroup
	resultCh := make(chan string, 0)
	stopCh := make(chan struct{})
	InterruptCh := make(chan os.Signal, 1)
	signal.Notify(InterruptCh, os.Interrupt)
	workNum := 10

	wg.Add(workNum)
	for i := 0; i < workNum; i++ {
		go func(i int) {
			defer wg.Done()
			for {
				select {
				case <-stopCh:
					fmt.Println("close:", i)
					return
				case resultCh <- GetData(i):
				default:
				}
			}
		}(i)
	}
	go func() {
		i := 0
		for item := range resultCh {
			i++
			fmt.Println("res:", item)
		}
	}()
	//专门写一个协程监听信号，然后关闭stopCh
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-InterruptCh:
				fmt.Println("receive close signal")
				close(stopCh)
				return
			default:
			}
		}
	}()
	wg.Wait() //当所有发送者结束后，需要关闭resultCh
	close(resultCh)
}
