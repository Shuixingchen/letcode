package myhttp

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func HttpClient() {
	resp, err := http.Get("https://www.baidu.com/")
	if err!= nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}
