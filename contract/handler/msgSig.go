package handler

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/crypto"
)

const (
	/*
	 * bytes4(keccak256('supportsInterface(bytes4)')) == 0x01ffc9a7
	 */
	_INTERFACE_ID_ERC165 = "0x01ffc9a7"

	// Equals to `bytes4(keccak256("onERC721Received(address,address,uint256,bytes)"))`
	// which can be also obtained as `IERC721Receiver(0).onERC721Received.selector`
	_ERC721_RECEIVED = "0x150b7a02"

	/*
	 *     bytes4(keccak256('balanceOf(address)')) == 0x70a08231
	 *     bytes4(keccak256('ownerOf(uint256)')) == 0x6352211e
	 *     bytes4(keccak256('approve(address,uint256)')) == 0x095ea7b3
	 *     bytes4(keccak256('getApproved(uint256)')) == 0x081812fc
	 *     bytes4(keccak256('setApprovalForAll(address,bool)')) == 0xa22cb465
	 *     bytes4(keccak256('isApprovedForAll(address,address)')) == 0xe985e9c5
	 *     bytes4(keccak256('transferFrom(address,address,uint256)')) == 0x23b872dd
	 *     bytes4(keccak256('safeTransferFrom(address,address,uint256)')) == 0x42842e0e
	 *     bytes4(keccak256('safeTransferFrom(address,address,uint256,bytes)')) == 0xb88d4fde
	 *     => 0x70a08231 ^ 0x6352211e ^ 0x095ea7b3 ^ 0x081812fc ^
	 *        0xa22cb465 ^ 0xe985e9c5 ^ 0x23b872dd ^ 0x42842e0e ^ 0xb88d4fde == 0x80ac58cd
	 */
	_INTERFACE_ID_ERC721 = "0x80ac58cd"

	/*
	 *     bytes4(keccak256('name()')) == 0x06fdde03
	 *     bytes4(keccak256('symbol()')) == 0x95d89b41
	 *     bytes4(keccak256('tokenURI(uint256)')) == 0xc87b56dd
	 *     => 0x06fdde03 ^ 0x95d89b41 ^ 0xc87b56dd == 0x5b5e139f
	 */
	_INTERFACE_ID_ERC721_METADATA = "0x5b5e139f"

	/*
	 *     bytes4(keccak256('totalSupply()')) == 0x18160ddd
	 *     bytes4(keccak256('tokenOfOwnerByIndex(address,uint256)')) == 0x2f745c59
	 *     bytes4(keccak256('tokenByIndex(uint256)')) == 0x4f6ccce7
	 *     => 0x18160ddd ^ 0x2f745c59 ^ 0x4f6ccce7 == 0x780e9d63
	 */
	_INTERFACE_ID_ERC721_ENUMERABLE = "0x780e9d63"
)

// 使用eth对消息进行签名，并且验证签名人

func VerifySig(signAddr, msg, sig string) bool {
	return false
}

func SigMessage(msg string) {
	msgHash := sha256.Sum256([]byte(msg))
	// 准备私钥
	pkeyb, err := hex.DecodeString(privateHex)
	if err != nil {
		log.Fatalln(err)
	}
	// 基于secp256k1的私钥
	pkey, err := crypto.ToECDSA(pkeyb)
	if err != nil {
		log.Fatalln(err)
	}
	// 签名,会得到65字节数据，
	sig, err := crypto.Sign(msgHash[:], pkey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("sig length:", len(sig))
	fmt.Println("sig hex:", hex.EncodeToString(sig))
}

// func GetToenInterfaceId() {
// 	bytes4(keccak256("balanceOf(address)")) ^
// 		bytes4(keccak256("ownerOf(uint256)")) ^
// 		bytes4(keccak256("approve(address,uint256)")) ^
// 		bytes4(keccak256("getApproved(uint256)")) ^
// 		bytes4(keccak256("setApprovalForAll(address,bool)")) ^
// 		bytes4(keccak256("isApprovedForAll(address,address)")) ^
// 		bytes4(keccak256("transferFrom(address,address,uint256)")) ^
// 		bytes4(keccak256("safeTransferFrom(address,address,uint256)")) ^
// 		bytes4(keccak256("safeTransferFrom(address,address,uint256,bytes)"))
// }

func GetTokenInterfaceId() {
	balanceOf := crypto.Keccak256Hash([]byte("balanceOf(address)"))
	ownerOf := crypto.Keccak256Hash([]byte("ownerOf(uint256)"))
	approve := crypto.Keccak256Hash([]byte("approve(address,uint256)"))
	getApproved := crypto.Keccak256Hash([]byte("getApproved(uint256)"))
	setApprovalForAll := crypto.Keccak256Hash([]byte("setApprovalForAll(address,bool)"))
	isApprovedForAll := crypto.Keccak256Hash([]byte("isApprovedForAll(address,address)"))
	transferFrom := crypto.Keccak256Hash([]byte("transferFrom(address,address,uint256)"))
	safeTransferFrom := crypto.Keccak256Hash([]byte("safeTransferFrom(address,address,uint256)"))
	safeTransferFromB := crypto.Keccak256Hash([]byte("safeTransferFrom(address,address,uint256,bytes)"))
	fmt.Println(balanceOf, ownerOf, approve, getApproved, setApprovalForAll, isApprovedForAll, transferFrom, safeTransferFrom, safeTransferFromB)
}
