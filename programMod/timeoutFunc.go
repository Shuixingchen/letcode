package programMod

import (
	"errors"
	"time"
)

//函数超时机制，超过时间函数还没有返回，就直接返回超时错误

func Myfunc() string {
	time.Sleep(1 * time.Second)
	return "success"
}

func Do(timer *time.Timer) (string, error) {
	Done := make(chan string, 0)
	go func() {
		result := Myfunc()
		Done <- result
	}()
	select {
	case res := <-Done:
		return res, nil
	case <-timer.C:
		return "", errors.New("timeout")
	}
}
