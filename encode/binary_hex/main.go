package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

//pack() 把数据转为字节数组
func IntToBytes(n int64) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{}) //先传空字节数组，写入的时候会自动扩充
	binary.Write(bytesBuffer, binary.BigEndian, n)
	gbyte := bytesBuffer.Bytes()
	return gbyte
}

//binary.Write()原理
func BigEndianIntToBytes(n int64, k int) []byte {
	bs := make([]byte, k)
	switch k {
	case 2:
		binary.BigEndian.PutUint16(bs, uint16(n))
	case 4:
		binary.BigEndian.PutUint32(bs, uint32(n))
	case 8:
		binary.BigEndian.PutUint64(bs, uint64(n))
	}
	return bs
}

//PutUint16()的原理
/*
以num = 300为例，计算机加载进去内存后，用两个字节保存 00000001 00101100(不同cup架构存放会不同，但是byte()会读取最低位的字节)
1.byte(),会把uint16读取一个低字节，所以byte(300)返回00101100，也就是44 属于uint16的低位
2.把uint16右移8位，得到00000000 00000001，再取一个低字节00000001,属于uint16的高位
大端序，字节数组从0开始放高位到地位，所以b[0] = 1, b[1]= 44
*/
func MyPutUint16(b []byte, num uint16) {
	_ = b[1]
	b[1] = byte(num)
	b[0] = byte(num >> 8)
}
func MyPutUint32(b []byte, num uint32) {
	_ = b[3]
	b[3] = byte(num)
	b[2] = byte(num >> 8)
	b[1] = byte(num >> 16)
	b[0] = byte(num >> 24)
}

//十六进制字符串转字节数组，两个字符->一个字节
func HexToBytes(hexstr string) []byte {
	b, err := hex.DecodeString(hexstr)
	if err != nil {
		panic(err)
	}
	return b
}

// 字节转十六进制字符
func BytesToHex(b []byte) string {
	return hex.EncodeToString(b)
}

//unpack() 把字节数组为整数, 字节数组为大端序
func ByteToInt(b []byte) int64 {
	bytesBuffer := bytes.NewBuffer(b)
	l := len(b)
	switch l {
	case 2:
		var num int16
		binary.Read(bytesBuffer, binary.BigEndian, &num)
		return int64(num)
	case 4:
		var num int32
		binary.Read(bytesBuffer, binary.BigEndian, &num)
		return int64(num)
	case 8:
		var num int64
		binary.Read(bytesBuffer, binary.BigEndian, &num)
		return num
	}
	return -1
}

func BigEndianBytesToInt(b []byte) int64 {
	switch len(b) {
	case 2:
		res := binary.BigEndian.Uint16(b)
		return int64(res)
	case 4:
		res := binary.BigEndian.Uint32(b)
		return int64(res)
	case 8:
		res := binary.BigEndian.Uint64(b)
		return int64(res)
	}
	return -1
}

/*
以b[1,44]为例子
1.首先把每个字节的整数转uint16，表示用两个字节保存
00000000 00000001 //b[0]
00000000 00101100 //b[1]

2. 因为是大端序，需要把高字结位b[0]移动到左边。
00000001 00000000  //b[0]

3.使用与运算|, b[0]|b[1]得到新的uint16
00000001 00000000 b[0]
00000000 00101100 b[1]
00000001 00101100 uint16
*/
func MyUint16(b []byte) uint16 {
	_ = b[1]
	return uint16(b[1]) | uint16(b[0])<<8
}

//同理，可以得出uint32
func MyUint32(b []byte) uint32 {
	_ = b[3]
	return uint32(b[0])<<24 | uint32(b[1])<<16 | uint32(b[2])<<8 | uint32(b[3])
}

func stringToByte(str string) {
	hex.DecodeString(str)
}

func main() {
	str := "0x000000000000000000000000ef0031812faca5e803b13fca55abe9836356066a"
	str = str[2:]
	fmt.Println(str)
	res := HexToBytes(str)
	fmt.Println(res)
}
