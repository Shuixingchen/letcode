package handler

import (
	"context"
	"database/sql"
	"fmt"
	"letcode/crawler/fetchers"
	"letcode/crawler/models"
	"letcode/crawler/tokens"
	"os"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

type EtherscanHandler struct {
	DB      *sql.DB
	Fetcher *fetchers.EtherscanFetcher
}

func NewEtherscanHandler() *EtherscanHandler {
	db := NewLoader()
	fetcher := fetchers.NewEtherscanFetcher()
	return &EtherscanHandler{
		DB:      db,
		Fetcher: fetcher,
	}
}

func NewLoader() *sql.DB {
	dsn := "root:123456@tcp(127.0.0.1:3306)/eth_parser?charset=utf8"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	return db
}

func (h *EtherscanHandler) Handle() {
	for i := 0; i < 20; i++ {
		tokenHashs := h.Fetcher.SendRequest(i)
		for _, hash := range tokenHashs {
			if h.IsExist(hash) {
				continue
			}
			fmt.Println(hash)
			token := tokens.QueryERC20(hash)
			h.Save(token)
		}
	}
}

func (h *EtherscanHandler) IsExist(tokenHash string) bool {
	query := "select token_address from tokens where token_address = '" + tokenHash + "'"
	_, err := models.GetTokenHash(context.Background(), query, h.DB)
	if err != nil {
		return false
	}
	return true
}

func (h *EtherscanHandler) Save(token *models.Token) {
	query := models.TokenSelectQueryPrefix + "where token_address = " + token.Addr
	_, err := models.GetToken(context.Background(), query, h.DB)
	tokens := make([]*models.Token, 1)
	tokens[0] = token
	if err != nil {
		data := models.TokenBatchInsertion(tokens)
		h.File(data)
	}
}

func (h *EtherscanHandler) File(data string) error {
	dir := "./output"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0777)
		if err != nil {
			log.Errorf("mkdir error: %v", err)
			return err
		}
	}
	filename := "token.csv"
	file, err := os.OpenFile(
		filename,
		os.O_WRONLY|os.O_CREATE|os.O_APPEND,
		0666,
	)
	_, err = file.Write([]byte(data))
	return err
}
