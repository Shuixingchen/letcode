package programMod

/*
管道模式
*/
//这个是数据源
func Echo(arr []int) <-chan int {
	ch := make(chan int)
	go func() {
		for _, i := range arr {
			ch <- i
		}
		close(ch) //如果写完了数据，就要关闭chan,否则for range 通道时会一直阻塞
	}()
	return ch
}

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

type EchoFunc func([]int) <-chan int
type PipeFunc func(<-chan int) <-chan int

//数据从echo出来。依次给PipeFunc处理
func Pipeline(nums []int, echo EchoFunc, pipeFns ...PipeFunc) <-chan int {
	ch := echo(nums)
	for i := range pipeFns {
		ch = pipeFns[i](ch)
	}
	return ch
}
