package main

import (
	"letcode/contract/handler"
)

func main() {
	// handler.SendTx()
	// handler.ParseAddressFromSigTx()
	// handler.QueryERC20("0x17eb0d548306372293C67A7DAD5c6bCBfE5593F8")
	// handler.QueryERC721("0x40875223D61a688954263892d0f76c94fd6B3D4A") //erc1155
	// handler.AccountInfo()
	// handler.QueryERC20("0x4047DeF04Ae123Bf223808119B9685104A33D5f7", "0xdBe64A759Da9ac5c4bD4782585a9D3b9711eDfaD")
	// handler.QueryEventLog("0x4047DeF04Ae123Bf223808119B9685104A33D5f7", 26431955, 26431956)
	// handler.ListenEvent()
	// handler.ParseNFTTx("0x80c3b11a8bef99cee83e233b751d23764e895070b522c3baa50ebf6882a27767")
	// ShowSig()
	// handler.ParserTxInput("0xc7fec0f7041276d51be25786fb8dfdff3662e5da2866b9d7a15f039bb88381e7")
	// handler.ParserTxInput("0x33243cfe206bcaccba46488e7f53c84cb4ef01665e9613f62833533faec2bdfd") //eth
	// handler.ParserTxInput("0x8461106a76b3a4100eb76e3e614d9f09a99382cecf5eea43ef61cfa0f7c48c7a") //weth
	// handler.ParserTxInput("0x5b614add5a735e59de090f828e4c14077d9c66949336cdd5927e9d0aab7655b0") // transferbatch+eth
	// handler.ParserApi()
	// handler.ParserTxInput("0xc7fec0f7041276d51be25786fb8dfdff3662e5da2866b9d7a15f039bb88381e7")
	// handler.ParseBlock("0xadcd94d5096e1f679fa7cc5a84420b98de3bf4c9522340aeddb610176ffd84ae")
	// handler.IsNFTByClient("0x76280AF9D18a868a0aF3dcA95b57DDE816c1aaf2")
	// handler.ParseTx("0x9891a0d43fc755e80ba7c374dad45503b9b67fa50603f34d8fcc0142878c9080")
	// handler.QueryTokenURI("0x46C31b6B330c4522D6b37CB3CEcDF2dA9fF46F61", "113")
	// handler.QueryERC721("0x4abAdb072fBa668395d0cF39dc89326662592D19")
	// handler.QueryERC721Balance("0xB5C747561a185A146f83cFff25BdfD2455b31fF4", "4656")
	// handler.QueryERC20("0x65498bf901a0c47ba9507c8a778d2bdee4db12b4", "0x0a55bca59602db94ff7df975fe8dfda6e6c92ff6")
	// handler.QueryERC1155Balance("0xAd6Dc35442d766f87f9296F17BA45e23518Bc5F3", "0x497833ED78601b6b62B705b92De3784479511403", "1112")
	// handler.GoKit("0xc7fec0f7041276d51be25786fb8dfdff3662e5da2866b9d7a15f039bb88381e7") //没有opensea的transfer
	// handler.GoKit("0x33243cfe206bcaccba46488e7f53c84cb4ef01665e9613f62833533faec2bdfd") // opensea的sell
	// handler.GoKit("0x8461106a76b3a4100eb76e3e614d9f09a99382cecf5eea43ef61cfa0f7c48c7a") // 使用erc20-nft
	// handler.GetENSTokenID("fudging")
	// handler.ParseENSRegister("0x27ab7b42b7858f2f69f5985ecc31d7cd6a8eb6e66109bf8308562050b53bcb9d")
	// handler.SignMessage("sign my message")
	// handler.VerifySig("sign my message", "8b3f09cc8a49582c6ded832a8062a8f5bf66a701224e83d3c569681a0b6436de0bb22d3f21546c4c4b6a78f7c36322080d4e28a35791fbe80124396c85a5502a00")
	handler.VerifyHandler()
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
		"SetMintLimit(address,uint256)",
	}
	methods := []string{
		"transfer(address,uint256)",
		"setMintLimit(address,uint256)",
	}
	for _, event := range events {
		handler.Signature(event, true)
	}
	for _, method := range methods {
		handler.Signature(method, false)
	}
}

func QueryEventLog() {
	handler.QueryEventLog("0x2079812353e2c9409a788fbf5f383fa62ad85be8", 14338777, 14338777)
}

func ParseNFTTx() {
	// mainnet opensea nft转账交易
	handler.ParseNFTTx("0x33243cfe206bcaccba46488e7f53c84cb4ef01665e9613f62833533faec2bdfd")
}
