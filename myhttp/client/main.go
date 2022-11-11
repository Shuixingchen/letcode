package client

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"time"
)

func main() {
	//自定义request
	request, err := http.NewRequest("get", "https://xtz.getblock.io/mainnet/chains", http.NoBody)
	if err != nil {
		fmt.Println(err)
	}
	request.Header.Add("x-api-key", "0ee52bce-e6c2-4e3e-987b-629a3e1d4a08")
	request.Header.Add("Content-type", "application/json")

	//自定义transport
	var tlsConfig *tls.Config
	transport := &http.Transport{
		Dial:                  (&net.Dialer{}).Dial,
		TLSClientConfig:       tlsConfig,
		IdleConnTimeout:       60,
		ResponseHeaderTimeout: 60,
		ExpectContinueTimeout: 60,
		MaxIdleConns:          5,
		MaxIdleConnsPerHost:   5,
	}

	//自定义client
	client := http.Client{
		Transport: transport,
	}
	response, err := client.Do(request)
	fmt.Println(response)
}

// 设置请求time out
func RequestTime() {
	timer := time.NewTimer(5 * time.Second)
	req, _ := http.NewRequest("get", "https://xtz.getblock.io/mainnet/chains", http.NoBody)
	DoRequest(timer, req)
}
func DoRequest(timer *time.Timer, request *http.Request) (*http.Response, error) {
	type resType struct {
		response *http.Response
		err      error
	}
	done := make(chan resType, 0)
	go func() {
		var res resType
		res.response, res.err = http.DefaultClient.Do(request)
		done <- res
	}()
	select {
	case res := <-done:
		return res.response, res.err
	case <-timer.C:
		return nil, fmt.Errorf("request time out")
	}
}
