package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/panjf2000/ants"
)

func demoFunc() {
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Hello World!")
}

func main() {
	pool, _ := ants.NewPool(10)
	defer pool.Release()

	var wg sync.WaitGroup
	// 任务函数
	syncCalculateSum := func() {
		demoFunc()
		wg.Done()
	}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		_ = pool.Submit(syncCalculateSum)
	}
	wg.Wait()
	fmt.Printf("running goroutines: %d\n", ants.Running())
	fmt.Printf("finish all tasks.\n")
}
