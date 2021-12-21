package main

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

//logrus拥有六种日志级别：debug、info、warn、error、 fatal和panic
//前四种不会影响程序的执行，只会打印日志
//fatal：输出日志的同时，调用os.Exit(1)方法退出，小提示：defer不会执行
//panic:输出日志的同时，调用panic方法，但defer会执行

func InitLog(logLevel, path string) {
	//1.设置日志显示格式
	formatter := &log.TextFormatter{
		FullTimestamp: true,
	}
	log.SetFormatter(formatter)

	//2.设置日志等级
	level, err := log.ParseLevel(logLevel)
	if err != nil {
		log.Errorf("init logger error. %+v", errors.WithStack(err))
	}
	log.SetLevel(level)

	//3. 日志写入到io.write
	log.SetOutput(os.Stdout)
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(file)
	} else {
		fmt.Println(err)
	}
}

func main() {
	InitLog("debug", "logs/log.log")
	defer func() {
		fmt.Println("finish")
	}()
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")

	log.Warn("show warn error")
	log.Error("show error")
	// log.Fatal("show fatal error")
	// log.Panic("shwo panic error")

	fmt.Println("code finish")
}
