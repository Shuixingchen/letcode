package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

/*
websocket, 首先创建http服务，然后通过upgrader获取websocket连接
*/
var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func HandleWs(w http.ResponseWriter, r *http.Request) {
	var (
		wsConn *websocket.Conn
		err    error
	)
	wsConn, err = upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("upgrade faile:", err)
	}
	go SendMsg(wsConn)
	for {
		msgType, msg, err := wsConn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		if err = wsConn.WriteMessage(msgType, msg); err != nil {
			return
		}
	}
}

func SendMsg(wsConn *websocket.Conn) {
	ticker := time.NewTicker(time.Second * 3) // 运行时长
	var i int64
	for {
		select {
		case <-ticker.C:
			i++
			wsConn.WriteMessage(1, []byte("send:"+strconv.FormatInt(i, 10)))
		}
	}
}

func main() {
	http.HandleFunc("/", HandleWs)
	fmt.Printf("listen on:%d", 8080)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
