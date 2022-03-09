package handler

import (
	"context"
	"encoding/hex"
	"fmt"
	"letcode/contract/client"
	"log"
	"math/big"
	"strings"

	"github.com/btccom/gokit/explorer/ethereum/decode"
)

var (
	// common methods shared by token contracts
	transferSig     = new(big.Int).SetBytes([]byte{0xa9, 0x05, 0x9c, 0xbb})
	balanceOfSig    = new(big.Int).SetBytes([]byte{0x70, 0xa0, 0x82, 0x31})
	totalSupplySig  = new(big.Int).SetBytes([]byte{0x18, 0x16, 0x0d, 0xdd})
	transferFromSig = new(big.Int).SetBytes([]byte{0x23, 0xb8, 0x72, 0xdd})
	allowenceSig    = new(big.Int).SetBytes([]byte{0xdd, 0x62, 0xed, 0x3e})
	approveSig      = new(big.Int).SetBytes([]byte{0x09, 0x5e, 0xa7, 0xb3})

	// optional
	nameSig     = new(big.Int).SetBytes([]byte{0x06, 0xfd, 0xde, 0x03})
	symbolSig   = new(big.Int).SetBytes([]byte{0x95, 0xd8, 0x9b, 0x41})
	decimalsSig = new(big.Int).SetBytes([]byte{0x31, 0x3c, 0xe5, 0x67})

	// ERC721 specific
	ownerSig = new(big.Int).SetBytes([]byte{0x63, 0x52, 0x21, 0x1e})
)
var ERC20 = []*big.Int{
	transferSig,
	totalSupplySig,
	balanceOfSig,
	transferFromSig,
	allowenceSig,
	approveSig,
}

var ERC721 = []*big.Int{
	transferSig,
	balanceOfSig,
	approveSig,
}

var (
	erc20Addr = "0x022a22918a11dBD673EFa3619d648f53EAAe9355"
)

func IsToken(addr string) {
	ec := client.CreateEthClient()
	code, err := ec.GetCodeLatest(context.Background(), erc20Addr)
	if err != nil {
		log.Fatal(err)
	}
	code = strings.TrimPrefix(code, "0x")

	isERC721 := IsERC(code, ERC721)
	isERC20 := IsERC(code, ERC20)
	fmt.Println(isERC721, isERC20)
}

func IsERC(input string, erc []*big.Int) bool {
	code, err := hex.DecodeString(input)
	if err != nil {
		return false
	}
	instrs, err := decode.Disassemble(code)
	if err != nil {
		return false
	}
	sigs := decode.GetMethodSigs(instrs)
	return matchSigs(sigs, erc)
}

func matchSigs(sigs []*big.Int, targets []*big.Int) bool {
	for _, target := range targets {
		matched := false
		for _, sig := range sigs {
			if sig.Cmp(target) == 0 {
				matched = true
				break
			}
		}
		if matched == false {
			return false
		}
	}
	return true
}
