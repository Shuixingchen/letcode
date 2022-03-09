package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Item struct {
	ID       int
	Symbol   string
	Address  string
	ImageUrl string `json:"image_url"`
	Name     string
	Decimals int
	EthPrice string `json:"eth_price"`
	UsdPrice string `json:"usd_price"`
}

func main() {
	l := NewLoader()
	list := Crawler()
	l.Insert(list)
}

func Crawler() []Item {
	res, err := http.Get("https://api.opensea.io/tokens/?limit=1")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	list := make([]Item, 0)

	if err := json.Unmarshal(body, &list); err != nil {
		log.Fatal(err)
	}
	return list
}

type Loader struct {
	db  *sql.DB
	dsn string
}

func NewLoader() Loader {
	dsn := "root:123456@tcp(127.0.0.1:3306)/nft_opensea?charset=utf8"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	return Loader{db: db, dsn: dsn}
}

func (l *Loader) Insert(list []Item) {
	sqlInsert := "INSERT Into nft_ranking_list(nid, symbol, address, image_url, name,decimals,eth_price,used_price,stime)values(?, ?, ?, ?, ?,?,?,?,?)"
	stmt, err := l.db.Prepare(sqlInsert) //编译sql，并且返回一个stmt
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()
	for _, val := range list {
		_, err = stmt.Exec(val.ID, val.Symbol, val.Address, val.ImageUrl, val.Name, val.Decimals, val.EthPrice, val.UsdPrice, time.Now()) //给这个stmt传参并且执行
	}
}
