package main

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	Name string
	Id   int64
}

func DeferTime() {
	startedAt := time.Now()
	defer fmt.Println(time.Since(startedAt))

	time.Sleep(time.Second)
}
func DeferLog(data string) (user User) {
	defer func() {
		user.Name = "dd"
		fmt.Println(user)
	}()
	if len(data) > 0 {
		user.Name = "aa"
		user.Id = 1
		return
	}
	return
}

func DeferErr() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		} else {
			fmt.Println("r=nil")
		}
	}()
	panic(errors.New("aa"))
}

func DeferArgs() {
	user := []int{1, 2, 3}
	defer func() {
		user = append(user, 4)
	}()

}

func main() {
	DeferErr()
}
