package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

//gob包，需要一个io.ReadWrite参数，分别包装一层实现编码后写入，解码后读出。
/*
enc := gob.NewEncoder(conn), enc.Encode(data), 相当于编码后conn.write(data)，编码器写入
dec := gob.NewDecoder(conn), dec.Decode(&m) 解码到变量m，解码器读取
*/
//也可以使用json编解码
/*
enc := json.NewEncoder(conn), json.Encode(data), 相当于编码后conn.write(data)
dec := json.NewDecoder(conn), json.Decode(&m) 解码到变量m
*/

func main() {
	info := map[string]string{
		"name":    "C语言中文网",
		"website": "http://c.biancheng.net/golang/",
	}
	name := "demo.gob"
	EncodeToByte(name, info)
	DecodeFromFile(name)
}
func EncodeToByte(name string, data interface{}) {
	fd, _ := os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0777)
	defer fd.Close()
	enc := gob.NewEncoder(fd)
	if err := enc.Encode(data); err != nil {
		fmt.Println(err)
	}
}
func DecodeFromFile(name string) {
	var m map[string]string
	fd, _ := os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0777)
	defer fd.Close()
	D := gob.NewDecoder(fd) //生成解码器
	D.Decode(&m)
	fmt.Println(m)
}
