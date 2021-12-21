package main

import (
	"fmt"
	"unsafe"
)

//go 汇编 go tool compile -S main.go

var id int = 10

/*
内存对齐：cpu读取数据不是一个字节一个字节读的，而是8个字节（64位）,以空间换时间
规则1：变量的地址(迁移量)必须是变量类型大小的整数倍。

go检测内存对齐工具fieldalignment
安装后使用命令：fieldalignment main.go
*/
type Part1 struct {
	a bool  //1
	b int32 //4
	c int8  //1
	d int64 //8
	e byte  //1
}

func main() {
	part1 := Part1{}
	//part1内存对齐后  aXXX|bbbb|cXXX|XXXX|dddd|dddd|e
	fmt.Printf("part1 size: %d, align: %d\n", unsafe.Sizeof(part1), unsafe.Alignof(part1))
}
