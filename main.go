package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

type CKBProcessor struct {
	closed      chan struct{} // indicate if syncer is closed
	interruptCh chan os.Signal
	updateCh    chan uint64 // update chan used to receive new block number notification.
}

func main() {
	close := make(chan struct{})
	interCh := make(chan os.Signal)
	update := make(chan int)
	go MyFunc(close, update)
	ticker := time.NewTicker(1 * time.Second)
	signal.Notify(interCh, os.Interrupt)
	var i = 0
	for {
		select {
		case <-interCh:
			close <- struct{}{}
			fmt.Println("finish close")
			return
		case <-ticker.C:
			update <- i
			i++
		}
	}
}

func MyFunc(ch <-chan struct{}, update <-chan int) {
	for {
		select {
		case <-ch:
			fmt.Print("close")
			return
		case i := <-update:
			fmt.Println(i)
		}
	}
}
