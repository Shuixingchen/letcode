package main

import (
	"fmt"
	"io/ioutil"
	fbs "letcode/encode/flatbuff/fbs/block"
	"os"

	flatbuffers "github.com/google/flatbuffers/go"
)

type Block struct {
	Id   int64
	Hash string
	Flag bool
	Txs  []Tx
}
type Tx struct {
	Hash  string
	Value float64
}

func main() {
	var (
		filename = "./blok"
	)
	// data := EncodeToByte()
	// SaveToFile(filename, data)
	b := DecodeToBlock(filename)
	fmt.Println(b)
}

// 解码数据
func DecodeToBlock(filename string) Block {
	var (
		block Block
	)
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	//传入二进制数据
	b := fbs.GetRootAsBlock(buf, 0)
	block.Flag = b.Flag()
	block.Hash = string(b.Hash())
	block.Id = b.Id()
	len := b.TxsLength()
	for i := 0; i < len; i++ {
		tx := new(fbs.Tx)
		ntx := new(Tx)
		if b.Txs(tx, i) {
			ntx.Hash = string(tx.Hash())
			ntx.Value = tx.Value()
		}
		block.Txs = append(block.Txs, *ntx)
	}
	return block
}

func EncodeToByte() []byte {
	txone := Tx{Hash: "adfadf", Value: 123}
	txtwo := Tx{Hash: "adfadf", Value: 456}
	block := Block{
		Id:   1,
		Hash: "aadd",
		Flag: true,
	}
	//初始化buffer，大小为0，会自动扩容
	builder := flatbuffers.NewBuilder(0)
	//第一个交易
	txoneh := builder.CreateString(txone.Hash)
	fbs.TxStart(builder)
	fbs.TxAddHash(builder, txoneh)
	fbs.TxAddValue(builder, txone.Value)
	ntxone := fbs.TxEnd(builder)
	//第二个交易
	txtwoh := builder.CreateString(txtwo.Hash)
	fbs.TxStart(builder)
	fbs.TxAddHash(builder, txtwoh)
	fbs.TxAddValue(builder, txtwo.Value)
	ntxtwo := fbs.TxEnd(builder)

	//block
	//先处理数组，string等非标量
	fbs.BlockStartTxsVector(builder, 2)
	builder.PrependUOffsetT(ntxtwo)
	builder.PrependUOffsetT(ntxone)
	txs := builder.EndVector(2)
	bh := builder.CreateString(block.Hash)
	//再处理标量
	fbs.BlockStart(builder)
	fbs.BlockAddId(builder, block.Id)
	fbs.BlockAddHash(builder, bh)
	fbs.BlockAddFlag(builder, block.Flag)
	fbs.BlockAddTxs(builder, txs)
	nb := fbs.BlockEnd(builder)
	builder.Finish(nb)
	buf := builder.FinishedBytes() //返回[]byte
	return buf
}

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
func SaveToFile(filename string, data []byte) (int, error) {
	var (
		f   *os.File
		err error
	)
	if checkFileIsExist(filename) { //如果文件存在
		f, err = os.OpenFile(filename, os.O_APPEND, 0666) //打开文件
	} else {
		f, err = os.Create(filename) //创建文件
	}
	if err != nil {
		panic(err)
	}
	return f.Write(data)
}
