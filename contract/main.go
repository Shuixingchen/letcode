package main

import (
	"letcode/contract/handler"
)

func main() {
	// var wg sync.WaitGroup
	// wg.Add(1)
	// go func() {
	// 	handler.SubEvent()
	// 	wg.Done()
	// }()
	// time.Sleep(1 * time.Second)
	// handler.AccountInfo()
	// handler.QueryERC20("0x17eb0d548306372293C67A7DAD5c6bCBfE5593F8")
	// handler.QueryERC721("0xfF06b40b853b2700Afa5019aBE084469F10b63a5")
	// QueryEventLog()
	// ParseNFTTx()
	// ShowSig()
	handler.GoKit("0x33243cfe206bcaccba46488e7f53c84cb4ef01665e9613f62833533faec2bdfd")
	// SupportsInterface()
	// wg.Wait()
}

func ExecERC() {
	// handler.ExecERC20("0x17eb0d548306372293C67A7DAD5c6bCBfE5593F8", "0xdd132adb1a045ff37575a86734f153d8fcec90b1")

	handler.ExecERC721("0xfF06b40b853b2700Afa5019aBE084469F10b63a5", "0xe725D38CC421dF145fEFf6eB9Ec31602f95D8097", "http://www.baidu.com")
}

func ShowSig() {
	events := []string{
		"Transfer(address,address,uint256)",
		"Approval(address,address,uint256)",
		"ApprovalForAll(address,address,bool)",
		"TransferSingle(address,address,address,uint256,uint256)",
		"TransferBatch(address,address,address,uint256[],uint256[])",
		"OrdersMatched(bytes32,bytes32,address,address,uint256,bytes32)",
	}
	methods := []string{
		"transfer(address,uint256)",
	}
	for _, event := range events {
		handler.Signature(event, true)
	}
	for _, method := range methods {
		handler.Signature(method, false)
	}
}

func QueryEventLog() {
	handler.QueryEventLog("0x2079812353e2c9409a788fbf5f383fa62ad85be8", 14338777, 14338778)
}

func ParseNFTTx() {
	// mainnet opensea nft转账交易
	handler.ParseNFTTx("0x80c3b11a8bef99cee83e233b751d23764e895070b522c3baa50ebf6882a27767")
}

// 判断nft合约符合哪种协议
func SupportsInterface() {
	var (
		ERC165ID = [4]byte{0x01, 0xff, 0xc9, 0xa7}
	)
	handler.SupportsInterface("0xfF06b40b853b2700Afa5019aBE084469F10b63a5", ERC165ID)
}
