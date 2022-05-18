package handler

import (
	"database/sql"
	"encoding/hex"
	"fmt"
	fbs "letcode/encode/flatbuff/fbs/fbe"
)

// flatc -g fbs/explorer.fbs

type TxFlatBuffData struct {
	BlockNumber uint64
	BlockIndex  uint64
	Data        string
}

// 从数据库tx_flatbuff读取编码后的数据
func DecodeToTx(text string) *Transaction {
	s, _ := hex.DecodeString(text)
	txBin := fbs.GetRootAsTx(s, 0)
	var tx Transaction
	// inputs
	for i := 0; i < txBin.InputsLength(); i++ {
		var input Input
		inputBin := new(fbs.TxInput)
		if txBin.Inputs(inputBin, i) {
			input.ScriptAsm = hex.EncodeToString(inputBin.ScriptAsm())
			input.ScriptHex = hex.EncodeToString(inputBin.ScriptHex())
		}
		tx.Inputs = append(tx.Inputs, &input)
	}
	// outputs
	for i := 0; i < txBin.OutputsLength(); i++ {
		var output Output
		outputBin := new(fbs.TxOutput)
		if txBin.Outputs(outputBin, i) {
			output.ScriptAsm = hex.EncodeToString(outputBin.ScriptAsm())
			output.ScriptHex = hex.EncodeToString(outputBin.ScriptHex())
		}
		tx.Outputs = append(tx.Outputs, &output)
	}
	return &tx
}

func GetFlatBufBin() []*TxFlatBuffData {
	list := make([]*TxFlatBuffData, 0)
	dsn := "root:123456@tcp(127.0.0.1:3306)/ltc_parser?charset=utf8&maxAllowedPacket=17108864"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	rows, err := db.Query("select hex(UNCOMPRESS(fb_bin)) as bin,block_height,block_idx from tx_flatbuffer where block_height = %d and block_idx = %d", 1111, 0)
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var data TxFlatBuffData
		err := rows.Scan(&data.Data, &data.BlockNumber, &data.BlockIndex)
		if err != nil {
			fmt.Println(err)
		}
		list = append(list, &data)
	}
	return list
}
