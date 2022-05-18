package main

import (
	"letcode/crawler/handler"
)

func main() {
	handler := handler.NewEtherscanHandler()
	handler.Handle()
}
