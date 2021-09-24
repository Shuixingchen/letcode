package programMod

import (
	"errors"
	"time"
)

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
