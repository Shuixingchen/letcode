package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	arr := make([]string, 10)
	for i := 0; i < 10; i++ {
		arr[i] = strconv.Itoa(i)
	}
	for _, v := range arr {
		go func() {
			fmt.Printf("%s \n", v)
		}()
	}
	time.Sleep(time.Second * 1)
}
