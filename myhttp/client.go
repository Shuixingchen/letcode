package myhttp

import (
	"fmt"
	"net/http"
	"runtime"
)

func ClientRun() {
	for i:=0;i < 5;i++{
		resp,err := http.Get("https://baidu.com")
		if err != nil {
			fmt.Print("aa")
		}
		defer resp.Body.Close()
	}
	gnum := runtime.NumGoroutine()
	defer fmt.Print(gnum)
	//fmt.Print(resp)
}
