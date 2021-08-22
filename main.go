package main

import (
	"fmt"
	"math/big"
)

type HexBig struct {
	*big.Int
}

func main() {

	txFeeTotal := big.NewInt(16253653075803520)
	baseFeeBurn := big.NewInt(16253653075803520)
	blockTxFeeReward := new(big.Int)
	if enough := txFeeTotal.Cmp(baseFeeBurn); enough == 1 {
		blockTxFeeReward.Sub(txFeeTotal, baseFeeBurn)
	}
	TxFee := &HexBig{Int: blockTxFeeReward}
	fmt.Print(TxFee)
}
