package serve

import (
	"context"
	"errors"
	"fmt"
	"strconv"
)

type Response struct {
	Conn          *Conn
	HandlerHeader map[string]string
	ContentLength int64
	CancelCtx     context.CancelFunc
}

func (w *Response) Write(dataB []byte) (n int, err error) {
	lenData := len(dataB)
	if lenData == 0 {
		return 0, errors.New("length is 0")
	}
	w.ContentLength += int64(lenData)

	if dataB != nil {
		fmt.Println("start writing...")
		w.Conn.bufw.Write([]byte("HTTP/1.1 200 OK\n"))
		w.Conn.bufw.Write([]byte("Content-Type: application/json; charset=utf-8\n"))
		w.Conn.bufw.Write([]byte("Transfer-Encoding: chunked\n"))
		w.Conn.bufw.Write([]byte("Server: BTC.com\n"))
		w.Conn.bufw.Write([]byte("Content-Length: " + strconv.Itoa(lenData) + "\n"))
		w.Conn.bufw.Write(dataB)
		n = w.Conn.bufw.Available()
		err = w.Conn.bufw.Flush()
		fmt.Print(n, err)
		return
	}
	return 0, errors.New("dataB is nil")
}
