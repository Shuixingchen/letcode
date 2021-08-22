package main

import (
	"sync"

	log "github.com/sirupsen/logrus"
)

func main() {
	var res []int
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			res = append(res, i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	log.Info(len(res))
}
