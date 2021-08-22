package main

import (
	"fmt"
	"math/big"
	"strconv"
)

type EventLog struct {
	TimeStamp    int    `json:"timestamp,omitempty"`
	ContractAddr string `json:"address"`
	Name         string `json:"name"`
}

// WeiToEth converts wei to eth.
func WeiToEth(input *big.Int) float64 {
	var fi = new(big.Float)
	fi = fi.SetInt(input)
	fi = fi.Quo(fi, big.NewFloat(1e18))
	v, _ := fi.Float64()
	return v
}

func WeiToEth1(input *big.Int) string {
	// weiFloat := new(big.Float).SetInt(input)
	// ether := new(big.Float).Quo(weiFloat, big.NewFloat(1e18))
	// return ether.Text('f', 18)
	return ""
}
func WeiToGWei(input *big.Int) string {
	weiFloat := new(big.Float).SetInt(input)
	ether := new(big.Float).Quo(weiFloat, big.NewFloat(1e9))
	return ether.Text('f', 18)
}
func WeiToGWei1(input *big.Int) float64 {
	var fi = new(big.Float)
	fi = fi.SetInt(input)
	fi = fi.Quo(fi, big.NewFloat(1e9))
	v, _ := fi.Float64()
	return v
}
func turn() {
	baseFee := big.NewInt(31408309979)
	gasUsed := big.NewInt(10678917)
	burnFee := new(big.Int)
	burnFee.Mul(baseFee, gasUsed)
	res := WeiToEth1(burnFee)
	fmt.Print(res)
	f, err := strconv.ParseFloat(res, 64)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(f)
}

func main() {
	baseFee := big.NewInt(231408309979234233)
	res := WeiToGWei(baseFee)
	// last := fmt.Sprintf("%.18f", res)
	fmt.Print(res)
}
