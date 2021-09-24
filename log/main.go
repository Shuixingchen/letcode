package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

//logrus拥有六种日志级别：debug、info、warn、error、 fatal和panic
//前四种不会影响程序的执行，只会打印日志
//fatal：输出日志的同时，调用os.Exit(1)方法退出，小提示：如果函数下存在defer不会执行
//panic:输出日志的同时，调用panic方法，但defer会执行

func main() {
	defer func() {
		fmt.Println("finish")
	}()
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")

	log.Warn("show warn error")
	log.Error("show error")
	// log.Fatal("show fatal error")
	log.Panic("shwo panic error")

	fmt.Println("code finish")
}
