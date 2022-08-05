package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {
	rows := ReadFile()
	var wg sync.WaitGroup
	num, _ := strconv.Atoi(os.Args[1])
	fmt.Println("tcp num :", num)
	wg.Add(num)
	for i := 0; i < num; i++ {
		go start(rows[i][0], &wg)
	}
	wg.Wait()
}

func start(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	conn, err := net.Dial("tcp", "47.243.70.145:3333")
	if err != nil {
		fmt.Println("Dial error", err)
		return
	}
	fmt.Println("start " + name)
	cConnHandler(conn, strings.Trim(name, " "))
}

func cConnHandler(c net.Conn, name string) {
	var wg sync.WaitGroup
	wg.Add(2)
	go ReadLoop(c, name, &wg)
	go func() {
		c.Write([]byte("{\"id\":1,\"method\":\"mining.subscribe\",\"params\":[\"bmminer/0.0\"]}\n"))
		c.Write([]byte("{\"id\":2,\"method\":\"mining.authorize\",\"params\":[\"" + name + "\", \"\"]}\n"))
		wg.Done()
	}()
	wg.Wait()
	log.Println(name + " closed")
	c.Close()
}

func ReadLoop(c net.Conn, name string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		r := bufio.NewReader(c)
		line, _, err := r.ReadLine()
		if err != nil {
			if err != io.EOF {
				log.Println(name+" read error:", err)
			}
			break
		} else {
			log.Printf(name+" read %d bytes, content is %s\n", string(line))
		}
	}
}

func ReadFile() [][]string {
	// 绝对路径
	fs, err := os.Open("./bit.csv")
	if err != nil {
		fmt.Println(fmt.Sprintf("can not open the file, err is %+v", err))
	}
	defer fs.Close()

	// 读取所有
	r := csv.NewReader(fs)
	r.FieldsPerRecord = -1
	rows, _ := r.ReadAll()
	return rows
}
