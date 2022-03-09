package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	SqlExeHighLevel = "heigh"
	SqlMiddlehLevel = "middle"
	SqlLevelSep     = "."
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Sprintln(r)
		}
	}()
	retCh1 := make(chan string)
	go func() {
		for i := 0; i < 100; i++ {
			retCh1 <- strconv.Itoa(i)
			time.Sleep(3 * time.Second)
		}
	}()

	for {
		select {
		case ret := <-retCh1:
			fmt.Println(ret)
		case <-time.After(time.Second * 2):
			log.Panic("subscribe pending time out")
		}
	}
}
