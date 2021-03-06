package client

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
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
