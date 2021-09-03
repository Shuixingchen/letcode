package main

import (
	"encoding/binary"
	"fmt"
)

/*
go 转大小端序
Big-endian（大端序）： 数据的高位字节存放在地址的低端 低位字节存放在地址高端,符合我们查看的习惯
Little-endian（小端序）：数据的高位字节存放在地址的高端 低位字节存放在地址低端
*/

func main() {
	//00000001 00101100
	//1        44
	var aa uint16 = 300
	buf1 := make([]byte, 2)
	buf2 := make([]byte, 2)
	binary.BigEndian.PutUint16(buf1, aa)
	fmt.Println("大端序：", buf1)
	binary.LittleEndian.PutUint16(buf2, aa)
	fmt.Println("小端序：", buf2)
	fmt.Printf("第一个字节：%p,第二个字节：%p", &buf1[0], &buf2[1])
}
