package main

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

/*
计算机正常的内存增长方式是从低到高(当然栈不是)，取数据方式是从基址根据偏移找到他们的位置，
大端存储因为第一个字节就是高位，从而很容易知道它是正数还是负数，对于一些数值判断会很迅速.
小端序第一个字节是低位，很容易进行运算。
*/

/*
go 转大小端序
Big-endian（大端序）： 数据的高位字节存放在地址的低端 低位字节存放在地址高端,符合我们查看的习惯
Little-endian（小端序）：数据的高位字节存放在地址的高端 低位字节存放在地址低端
*/

func main() {
	byteToInt()
}

func byteToInt() {
	b, err := hex.DecodeString("008e060000000000")
	if err != nil {
		fmt.Println(err)
	}
	num := binary.LittleEndian.Uint64(b)
	fmt.Println(b, num)
}
