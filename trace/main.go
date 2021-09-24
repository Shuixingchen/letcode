package main

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		DoFun()
	}()
	ticker := time.NewTicker(1 * time.Second)
	for range ticker.C {
		gnum := runtime.NumGoroutine()
		fmt.Println("aaaa" + strconv.Itoa(gnum))

	}
	
}
func DoFun() {
	defer func() {
		if r := recover(); r != nil {
			log.Error("recovered from critical issue: ")
		}
	}()
	for i := 0; i < 1000; i++ {
		time.Sleep(1 * time.Second)
		if i == 3 {
			log.Panic("qqq")
			// log.WithFields(log.Fields{
			// 	"error": i,
			// }).Fatal("dofunc exit")
		}
	}
}
