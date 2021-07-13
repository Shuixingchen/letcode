package main

import "fmt"

type Block struct {
	a int
}
type Demo interface{
	call(int)
}
func(b *Block)call(num int){
	b.a = 11
}

func List(d Demo){
	fmt.Println(d)
}
func main() {
	block := Block{
		1,
	}
	block.call(1)

}