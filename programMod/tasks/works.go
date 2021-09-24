package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"sync"
	"syscall"
	"time"
)

//多协程跑脚本任务

type Handler struct {
	TaskCh      chan int
	ResultCh    chan string
	Start       int
	End         int
	wg          sync.WaitGroup
	WorkerNum   int
	InterruptCh chan os.Signal //接受os信号
	IsClose     bool
}

func main() {
	handler := NewHandler()
	handler.Init()
	handler.Serve()
}

func NewHandler() *Handler {
	return &Handler{
		TaskCh:      make(chan int),
		ResultCh:    make(chan string),
		Start:       0,
		End:         10,
		WorkerNum:   5,
		wg:          sync.WaitGroup{},
		InterruptCh: make(chan os.Signal),
		IsClose:     false,
	}
}
func (h *Handler) Init() {
	runtime.GOMAXPROCS(256)
	if h.Start > h.End {
		panic("start > End")
	}
	signal.Notify(h.InterruptCh, os.Interrupt, syscall.SIGTERM) //接受系统中断信号，会写入通道
}

func (h *Handler) Serve() {
	//1.把任务写入队列
	go func(s, e int) {
		defer close(h.TaskCh) //1.1 任务全部写入后及时关闭通道
		for i := s; i <= e; i++ {
			if h.IsClose {
				return
			}
			h.TaskCh <- i
		}
	}(h.Start, h.End)

	//2.创建worker,每个worker循环监听任务队列
	h.wg.Add(h.WorkerNum)
	for j := 0; j < h.WorkerNum; j++ {
		go h.StartWork()
	}

	//3.返回结果，循环读取结果队列
	go func() {
		for res := range h.ResultCh {
			time.Sleep(1e9)
			fmt.Println("res:", res)
		}
		for {
			res := <-h.ResultCh
			time.Sleep(1e9)
			fmt.Println("res:", res)
		}
	}()
	//4.监听系统信号，提前关闭任务
	go func() {
		<-h.InterruptCh
		fmt.Print("receive os signal closing \n")
		h.Close()
	}()
	//wait是等待worker协程，worker完成后，结果通道也就没有数据写了，要及时关掉
	h.wg.Wait()
	close(h.ResultCh)

}

//worker是死循环，任务读完或者有中断信号要退出
func (h *Handler) StartWork() {
	defer func() {
		//time.Sleep(2 * time.Second) //延迟1s让结果处理协程处理
		h.wg.Done()
	}()
	for {
		select {
		case t, ok := <-h.TaskCh:
			if ok == false {
				return
			}
			res := GetData(t)
			fmt.Println("input result:", res)
			h.ResultCh <- res
		default:
		}
	}
}

func (h *Handler) Close() {
	h.IsClose = true
}

func GetData(number int) string {
	return strconv.Itoa(number)
}
