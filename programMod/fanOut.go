package programMod

import (
	"strconv"
	"sync"
)

/*
需求：一组user,只有Name的值，需要远程请求到Score的值
*/

//开启指定个协程处理数组，使用扇出方式处理
type User struct {
	Name  string
	Score int64
}

//模拟远程调用数据
func Dodata(user *User) *User {
	user.Score = int64(len(user.Name))
	return user
}

func main() {
	s := CreatData()
	var wg sync.WaitGroup
	outch := make(chan *User, 0)
	wg.Add(10)
	Handler(outch, 10, &wg, s, Dodata)

	wg.Wait() //worker处理完需要及时关闭outch
	close(outch)
}

func Handler(outch chan *User, number int, wg *sync.WaitGroup, s []*User, workerFun func(*User) *User) {
	inch := make(chan *User, 0)
	//协程1：把需要处理的参数写入inch
	go func() {
		for _, item := range s {
			inch <- item
		}
		close(inch)
	}()
	//协程2：开启number个协程，同时读取参数，把结果写入outch
	for i := 0; i < number; i++ {
		go func() {
			defer wg.Done()
			for item := range inch {
				res := workerFun(item)
				outch <- res
			}
		}()
	}
}

func CreatData() []*User {
	var list []*User
	for i := 0; i < 10; i++ {
		u := User{Name: strconv.Itoa(i)}
		list = append(list, &u)
	}
	return list
}
