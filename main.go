package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
)

type Response struct {
	Success bool `json:"success"`
	Data    List `json:"data"`
}
type List struct {
	Begin *big.Int `json:"begin"`
	Tx    []TxData `json:"tx"`
}
type TxData struct {
	HashNext     bool    `json:"hashNext"`
	Total        int     `json:"total"`
	Transactions []SolTx `json:"transactions"`
}
type SolTx struct {
	ID        string   `json:"_id"`
	Src       string   `json:"src"`
	Dst       string   `json:"dst"`
	Lamport   *big.Int `json:"lamport"`
	BlockTIme *big.Int `json:"blockTime"`
	Slot      *big.Int `json:"slot"`
	TxHash    string   `json:"txHash"`
	Fee       *big.Int `json:"fee"`
}

func main() {
	url := "https://api.solscan.io/account/soltransfer/txs?address=STEPNq2UGeGSzCyGVr2nMQAzf8xuejwqebd84wcksCK&offset=10000&limit=10"
	res, err := httpGet(url)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}

func httpGet(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		log.Println("get error")
	}
	defer response.Body.Close()
	body, err2 := ioutil.ReadAll(response.Body)
	if err2 != nil {
		log.Println("ioutil read error")
	}
	// var res Response
	// err = json.Unmarshal(body, &res)
	return string(body), err
	// return res, err
}
