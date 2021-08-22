package main

import (
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type UnconfirTxStat struct {
	FeeRate            []byte `db:"fee_rate"`
	TxSize             int64  `db:"tx_size"`
	TxSizeCount        int64  `db:"tx_size_count"`
	TxSizeDivMaxSize   []byte `db:"tx_size_divide_max_size"`
	TxDurationTimeRate []byte `db:"tx_duration_time_rate"`
	UpdateTime         int64  `db:"update_time"`
	CurrentBestFee     []byte `db:"current_best_fee"`
}

func main() {
	db, err := sqlx.Connect("mysql", "root:123456@tcp(127.0.0.1:3306)/forge?charset=utf8")
	sql := "SELECT fee_rate,tx_size,tx_size_count,tx_size_divide_max_size,tx_duration_time_rate,update_time,current_best_fee " +
		"FROM unconfirmed_tx_stats WHERE coin = ? "

	var result []*UnconfirTxStat
	var arr, newArr []string
	err = db.Select(&result, sql, "btc")
	if err != nil {
		fmt.Println(err)
	}
	for _, val := range result {
		arr = append(arr, string(val.TxDurationTimeRate))
	}
	str, _ := json.Marshal(arr)
	json.Unmarshal(str, &newArr)
}
