package handler

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"letcode/contract/artificial/erc20"
	"letcode/contract/artificial/erc721"
	"letcode/contract/client"
	"letcode/contract/models"
	"math/big"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"

	gokit "github.com/btccom/gokit/explorer/ethereum/types"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	ec  *ethclient.Client
	ecw *ethclient.Client
)
var (
	transferEventSig  = "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef" //Transfer(address,address,uint256)
	transferSingleSig = "0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62"
	transferBatchSig  = "0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb"
	ordersMatchedSig  = "0xc4109843e0b7d514e4c093114b863f8e7d8d9a458c372cd51bfe526b588006c9"
	approvalSig       = "0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"
	approvalForAllSig = "0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31"
)

type NFTTransaction struct {
	TxHash      string   `json:"transactionHash"`
	BlockNumber int64    `json:"blockNumber"`
	TimeStamp   int64    `json:"timestamp"`
	TokenAddr   string   `json:"tokenAddress"`
	Sender      string   `json:"sender"`
	Receiver    string   `json:"receiver"`
	Operator    string   `json:"operator"`
	TokenId     string   `json:"tokenId"`
	Amount      *big.Int `json:"amount"`
	Price       *big.Int
	PriceUnit   string
	PriceFrom   string
	Value       *big.Int
	txType      int8
}

// nft 转移
type TransferEvent struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
}

// erc1155
type TransferSingleEvent struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
}
type TransferBatchEvent struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
}

// opensea 交易
type OrdersMatchedEvent struct {
	BuyHash  [32]byte
	SellHash [32]byte
	Maker    common.Address
	Taker    common.Address
	Price    *big.Int
	Metadata [32]byte
}

// approve
type ApprovalEvent struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
}

// ApproveForAll
type ApprovalForAllEvent struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
}

func init() {
	var err error
	ec, err = ethclient.Dial("https://mainnet.infura.io/v3/40b043c639b44d72966d3535d523a4b3")
	if err != nil {
		log.Fatal(err)
	}
	ecw, err = ethclient.Dial("wss://ropsten.infura.io/ws/v3/40b043c639b44d72966d3535d523a4b3")
	if err != nil {
		log.Fatal(err)
	}
}

func AccountInfo() {
	account := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
	balance, err := ec.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(balance) // 25893180161173005034
}

// 与erc20合约交互
func QueryERC20(addr string) {
	contractAddr := common.HexToAddress(addr)
	var token = new(models.Token)
	var err error
	// 加载智能合约
	tc, err := erc20.NewErc20(contractAddr, ec)
	if err != nil {
		log.Fatal(err)
	}

	token.Addr = addr
	token.Name, err = tc.Name(nil)
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("unable to get the erc 20 token name")
	}
	token.Symbol, err = tc.Symbol(nil)
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("unable to get the symbol of the erc20 token")
	}
	decimals, err := tc.Decimals(nil)
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("unable to get the decimals of the erc 20 token")
	}
	token.Decimals = int64(decimals)
	supply, err := tc.TotalSupply(nil)
	if err != nil || supply == nil {
		log.WithFields(log.Fields{"err": err}).Error("unable to get the total supply of the erc 20 token")
	}
	if supply == nil {
		supply = new(big.Int)
	}
	token.InitTotalSupply = supply
	fmt.Println(token)
}

// 与erc721合约交互
func QueryERC721(addr string) {
	contractAddr := common.HexToAddress(addr)
	var token = new(models.Token)
	var err error
	// 加载智能合约
	tc, err := erc721.NewErc721(contractAddr, ec)
	if err != nil {
		log.Fatal(err)
	}
	token.Symbol, err = tc.Symbol(nil)
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("unable to get the symbol of the erc721 token")
	}
	token.Name, err = tc.Name(nil)
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("unable to get the erc721 token name")
	}
	// 查询NFT
	tokenId := big.NewInt(1)
	ownerAddr, err := tc.OwnerOf(nil, tokenId)
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("unable to get tokenId ownerof")
	}
	fmt.Println("NFT id=2 owner is:" + ownerAddr.Hex())
	tokenURI, err := tc.TokenURI(nil, tokenId)
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("unable to get tokenId ownerof")
	}
	fmt.Println("NFT id=2 uri is:" + tokenURI)

	fmt.Println(token)
}

// 执行合约方法
func ExecERC20(addr string, to string) {
	contractAddr := common.HexToAddress(addr)
	receiver := common.HexToAddress(to)
	// 加载智能合约
	tc, err := erc20.NewErc20(contractAddr, ec)
	if err != nil {
		log.Fatal(err)
	}
	txOpts := GetTxOpts()
	value := big.NewInt(10000000)
	tx, err := tc.Transfer(txOpts, receiver, value)
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("transfer")
	}
	fmt.Printf("tx sent: %s\n", tx.Hash().Hex())
}
func ExecERC721(addr, to, tokenUrl string) {
	contractAddr := common.HexToAddress(addr)
	receiver := common.HexToAddress(to)
	tc, err := erc721.NewErc721(contractAddr, ec)
	if err != nil {
		log.Fatal(err)
	}
	txOpts := GetTxOpts()
	// 铸造NFT, 不能立即知道铸造的nft的id,只能监听合约的铸造事件获取函数执行结果
	tx, err := tc.AwardItem(txOpts, receiver, tokenUrl)
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("transfer")
	}
	fmt.Printf("tx sent: %s\n", tx.Hash().Hex())
}

// 构造交易选项
func GetTxOpts() *bind.TransactOpts {
	privateKey, err := crypto.HexToECDSA("19935d89cb5c67657c64a6383d601e30f04eb179a0369227403e5343bba22107")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 获取当前地址的noce,当前的gasPrice
	nonce, err := ec.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	gasPrice, err := ec.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// 新建一个keyed transactor
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice
	return auth
}

// 订阅事件,客户端需要ws连接
func SubEvent() {
	contractAddress := common.HexToAddress("0xfF06b40b853b2700Afa5019aBE084469F10b63a5")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}
	logs := make(chan types.Log)
	// 订阅指定的log
	sub, err := ecw.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}
	ec.FilterLogs(context.Background(), query)
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog) // pointer to event log
		}
	}
}

// 查询event log
func QueryEventLog(addr string, start, end uint64) {
	// var txSingleLog TransferSingleEvent
	contractAddress := common.HexToAddress(addr)
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(start)),
		ToBlock:   big.NewInt(int64(end)),
		Addresses: []common.Address{
			contractAddress,
		},
	}
	logs, err := ec.FilterLogs(context.Background(), query)
	if err != nil {
		log.WithFields(log.Fields{"method": "FilterLogs"}).Error(err)
	}
	abi, err := abi.JSON(strings.NewReader(erc721.Erc721ABI))
	if err != nil {
		log.WithFields(log.Fields{"method": "abi.JSON"}).Error(err)
	}
	for _, log := range logs {
		switch log.Topics[0].Hex() {
		case transferEventSig:
			fmt.Printf("Log Name: Transfer\n")
			var transferEvent TransferEvent
			err := abi.UnpackIntoInterface(&transferEvent, "Transfer", log.Data) //data只包含未index的参数
			if err != nil {
				fmt.Println(err)
			}
			transferEvent.From = common.HexToAddress(log.Topics[1].String())

		case transferSingleSig:
			fmt.Printf("Log Name: TransferSingle\n")
			var transferEvent TransferSingleEvent
			err := abi.UnpackIntoInterface(&transferEvent, "TransferSingle", log.Data)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(transferEvent)
		case transferBatchSig:
			fmt.Printf("Log Name: TransferBatchSig\n")
		case ordersMatchedSig:
			fmt.Printf("Log Name: OrdersMatched\n")
			var transferEvent OrdersMatchedEvent
			err := abi.UnpackIntoInterface(&transferEvent, "OrdersMatched", log.Data)
			if err != nil {
				fmt.Println(err)
			}
			transferEvent.Maker = common.HexToAddress(log.Topics[1].String())
			transferEvent.Taker = common.HexToAddress(log.Topics[2].String())
		}
	}
}

// erc165
func SupportsInterface(addr string, interfaceId [4]byte) {
	contractAddr := common.HexToAddress(addr)
	// 加载智能合约
	tc, err := erc721.NewErc721(contractAddr, ec)
	if err != nil {
		log.Fatal(err)
	}

	isSupport, err := tc.SupportsInterface(nil, interfaceId)
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("SupportsInterface")
	}
	fmt.Println("is support", isSupport)
}

// 解析NFT交易，openSea
func ParseNFTTx(txHash string) {
	hash := common.HexToHash(txHash)
	_, isPending, err := ec.TransactionByHash(context.Background(), hash)
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("TransactionByHash")
	}
	if isPending {
		log.WithFields(log.Fields{"isPending": isPending}).Error("TransactionByHash")
	}
	txReceipt, err := ec.TransactionReceipt(context.Background(), hash)
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("TransactionReceipt")
	}
	for _, el := range txReceipt.Logs {
		if len(el.Topics) < 4 {
			continue
		}
		abi, err := abi.JSON(strings.NewReader(erc721.Erc721ABI))
		if err != nil {
			log.WithFields(log.Fields{"method": "abi.JSON"}).Error(err)
		}
		switch el.Topics[0].Hex() {
		case transferEventSig:
			fmt.Printf("Log Name: Transfer\n")
			var transferEvent TransferEvent
			err := abi.UnpackIntoInterface(&transferEvent, "Transfer", el.Data) //data只包含未index的参数
			if err != nil {
				fmt.Println(err)
			}
			transferEvent.From = common.HexToAddress(el.Topics[1].String())
			transferEvent.To = common.HexToAddress(el.Topics[2].String())
			transferEvent.TokenId = el.Topics[3].Big()
		case transferSingleSig:
			fmt.Printf("Log Name: TransferSingle\n")
			var transferEvent TransferSingleEvent
			err := abi.UnpackIntoInterface(&transferEvent, "TransferSingle", el.Data)
			if err != nil {
				fmt.Println(err)
			}
			transferEvent.Operator = common.HexToAddress(el.Topics[1].String())
			transferEvent.From = common.HexToAddress(el.Topics[2].String())
			transferEvent.To = common.HexToAddress(el.Topics[3].String())
		case transferBatchSig:
			fmt.Printf("Log Name: TransferBatchSig\n")
			var transferEvent TransferBatchEvent
			err := abi.UnpackIntoInterface(&transferEvent, "TransferBatch", el.Data)
			if err != nil {
				fmt.Println(err)
			}
			transferEvent.Operator = common.HexToAddress(el.Topics[1].String())
			transferEvent.From = common.HexToAddress(el.Topics[2].String())
			transferEvent.To = common.HexToAddress(el.Topics[3].String())
		case ordersMatchedSig:
			fmt.Printf("Log Name: OrdersMatched\n")
			var transferEvent OrdersMatchedEvent
			err := abi.UnpackIntoInterface(&transferEvent, "OrdersMatched", el.Data)
			if err != nil {
				fmt.Println(err)
			}
			transferEvent.Maker = common.HexToAddress(el.Topics[1].String())
			transferEvent.Taker = common.HexToAddress(el.Topics[2].String())
			transferEvent.Metadata = el.Topics[3]
		case approvalSig:
			fmt.Printf("Log Name: approval\n")
			var transferEvent ApprovalEvent
			err := abi.UnpackIntoInterface(&transferEvent, "TransferBatch", el.Data)
			if err != nil {
				fmt.Println(err)
			}
			transferEvent.Owner = common.HexToAddress(el.Topics[1].String())
			transferEvent.Approved = common.HexToAddress(el.Topics[2].String())
			transferEvent.TokenId = el.Topics[3].Big()
		case approvalForAllSig:
			fmt.Printf("Log Name: approval\n")
		}
	}
}

// 函数签名 对函数名Keccak-256(SHA-3)hash后取前4个字节,转十六进制字符
// 事件签名 对函数名Keccak-256(SHA-3)hash
func Signature(method string, isEvent bool) {
	hash := crypto.Keccak256Hash([]byte(method))
	hashHex := hash.Hex()
	if isEvent {
		fmt.Println("event " + method + ":" + hashHex)
		return
	}
	fmt.Println(method + ":" + hashHex[0:10]) //"transfer(address,uint256):0xa9059cbb"
}

func ShowERC20MethodSignature(erc20Addr string) {
	erc20Methers := []string{"0xa9059cbb", "0x70a08231", "0x60806040"}
	kec := client.CreateEthClient()
	code, err := kec.GetCodeLatest(context.Background(), erc20Addr)
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("GetCodeLatest")
	}
	for _, ms := range erc20Methers {
		b := strings.Contains(code, ms)
		if !b {
			log.WithFields(log.Fields{"method": ms}).Error("Contains")
		}
	}
}

type Receipt struct {
	BlockHash         string `json:"blockHash"`
	BlockNumber       string `json:"blockNumber"`
	ContractAddress   string `json:"contractAddress"`
	CumulativeGasUsed string `json:"cumulativeGasUsed"`
	EffectiveGasPrice string `json:"effectiveGasPrice,omitempty"`
	GasUsed           string `json:"gasUsed"`
	Status            string `json:"status"`
	TransactionHash   string `json:"transactionHash"`
	TransactionIndex  string `json:"transactionIndex"`
	Logs              []Log  `json:"logs"`
}

type Log struct {
	Address     string   `json:"address"`
	BlockHash   string   `json:"blockHash"`
	BlockNumber string   `json:"blockNumber"`
	Data        string   `json:"data"`
	LogIndex    string   `json:"logIndex"`
	Removed     bool     `json:"removed"`
	Topics      []string `json:"topics"`
	TxHash      string   `json:"transactionHash"`
	TxIndex     string   `json:"transactionIndex"`
	// custom fields
	TimeStamp string `json:"timestamp,omitempty"`
}

func (log *Log) ToTypesLog() *gokit.EventLog {
	eventLog := &gokit.EventLog{
		TxHash:       log.TxHash,
		ContractAddr: log.Address,
		Topics:       TurnToTypeTopic(log.Topics),
		Data:         log.Data,
		BlockHash:    log.BlockHash,
		Removed:      log.Removed,
	}
	eventLog.BlockNumber, _ = gokit.HexStringToHexUint64(log.BlockNumber)
	eventLog.TimeStamp, _ = gokit.HexStringToHexUint64(log.TimeStamp)
	eventLog.TxIndex, _ = gokit.HexStringToHexUint64(log.TxIndex)
	eventLog.LogIndex, _ = gokit.HexStringToHexUint64(log.LogIndex)
	return eventLog
}

type rpcResponse struct {
	ID      int             `json:"id"`
	JSONRPC string          `json:"jsonrpc"`
	Result  json.RawMessage `json:"result"`
	Error   interface{}     `json:"error"`
}

func TurnToTypeLogs(logs []Log) []*gokit.EventLog {
	eventLogs := make([]*gokit.EventLog, len(logs))
	for key := range logs {
		item := logs[key]
		eventLogs[key] = item.ToTypesLog()
	}
	return eventLogs
}
func TurnToTypeTopic(topics []string) []gokit.Topic {
	t := make([]gokit.Topic, len(topics))
	for key, item := range topics {
		t[key] = gokit.Topic(item)
	}
	return t
}

func GoKit(txHash string) {
	jsonStr := []byte(`{"jsonrpc":"2.0","method":"eth_getTransactionReceipt","params":["` + txHash + `"],"id":1}`)
	res, err := http.Post("https://mainnet.infura.io/v3/40b043c639b44d72966d3535d523a4b3", "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	var response rpcResponse
	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Println(err)
	}
	result := response.Result
	var receipt Receipt
	if err := json.Unmarshal([]byte(result), &receipt); err != nil {
		log.WithFields(log.Fields{"method": "GetTransactionReceipt", "params": result}).Panic(err)
	}

	gokit.ParseLog(TurnToTypeLogs(receipt.Logs))
}
