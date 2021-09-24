package main

import "fmt"

/*
管道模式
1.PipeFunc 管道函数，管道函数接收一个管道，返回一个新管道
2.Pipeline 参数，1数据源通道，2管道参数
*/

type PipeFunc func(<-chan int) <-chan int

func Sq(ch <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for val := range ch {
			out <- val * val
		}
		close(out)
	}()
	return out
}

func Odd(ch <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for val := range ch {
			if val%2 != 0 {
				out <- val
			}
		}
		close(out)
	}()
	return out
}

func Sum(ch <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		sum := 0
		for val := range ch {
			sum += val
		}
		out <- sum
		close(out)
	}()
	return out
}

//参数：1.数据源 2.管道函数
func Pipeline(dataSource <-chan int, pipeFns ...PipeFunc) <-chan int {
	for i := range pipeFns {
		dataSource = pipeFns[i](dataSource) //每次都是新的管道重新赋值给dataSource
	}
	return dataSource
}

//使用
func main() {
	data := []int{1, 2, 3, 4, 5, 6, 6, 7, 8, 8, 9, 9, 0}
	dataSource := make(chan int)
	go func() {
		for _, item := range data {
			dataSource <- item
		}
		close(dataSource)
	}()
	resCh := Pipeline(dataSource, Sq, Odd, Sum) //返回的是最后结果的管道

	for v := range resCh {
		fmt.Println("result: ", v)
	}
}
