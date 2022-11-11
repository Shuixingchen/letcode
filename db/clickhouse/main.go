package main

import (
	"database/sql"
	"fmt"
	"letcode/db/clickhouse/models"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
)

type ClickHouse struct {
	Conn *sql.DB
}

func NewClickHouse() *ClickHouse {
	var c ClickHouse
	c.Conn = clickhouse.OpenDB(&clickhouse.Options{
		Addr: []string{"54.169.100.47:9000"},
		Auth: clickhouse.Auth{
			Database: "test",
			Username: "inuser",
			Password: "meiyoumima",
		},
		Settings: clickhouse.Settings{
			"max_execution_time": 60,
		},
		DialTimeout: 5 * time.Second,
		Debug:       true,
	})
	c.Conn.SetMaxIdleConns(5)
	c.Conn.SetMaxOpenConns(10)
	c.Conn.SetConnMaxLifetime(time.Hour)
	return &c
}

const (
	createTable = `CREATE TABLE t_stock
	(
		"id" UInt32,
		"sku" String,
		"total" Decimal(16,2),
		"create_time" Datetime,    
	)ENGINE = MergeTree 
	ORDER BY (id, sku)
	PARTITION BY (create_time);`
)

type Stock struct {
	ID         uint64
	Sku        string
	Total      string
	CreateTime string
}

func main() {
	QueryTokenHolder()
	// Query()
}

func InsertTokenHolderBatch() {
	c := NewClickHouse()
	tokenHolders := models.GenerateTokenHolders(10000)
	insertSql := models.InsertTokenHolder(tokenHolders)
	// fmt.Println(insertSql)
	c.Exec(insertSql)
}

func QueryTokenHolder() {
	sql := "SELECT * FROM test.t_token_holder where token_addr = '0x28de56db50a1c9310f4eb478927dcc471f90a95c' and token_balance > 0"
	// sql := "SELECT * FROM test.t_token_holder where holder_addr = '0x000023ecf1665059624598912fe891b016dfaf64' and token_balance > 0"
	c := NewClickHouse()
	tokens := models.Query(c.Conn, sql)
	fmt.Println("count:", len(tokens))
}

func Query() {
	// sql := "SELECT count(1) FROM test.t_token_holder where token_addr = '0x28de56db50a1c9310f4eb478927dcc471f90a95c' and token_balance > 0"
	// sql_token_addr_count := "select count(DISTINCT token_addr) from t_token_holder"
	sql_holder_addr_count := "select count(DISTINCT holder_addr) from t_token_holder"
	c := NewClickHouse()
	var count uint64
	c.Conn.QueryRow(sql_holder_addr_count).Scan(&count)
	fmt.Println(count)
}

func (c *ClickHouse) InsertBatch() {
	sql := `INSERT INTO t_stock (id, sku, total, create_time) VALUES (1, 'aa', 11, '2022-01-01 12:00:00'), 
	(2, 'bb', 22, '2022-01-02 12:00:00')`
	c.Exec(sql)
}

func (c *ClickHouse) ShowTables() {
}

func (c *ClickHouse) CreateDB() (sql.Result, error) {
	sql := `CREATE DATABASE  test`
	return c.Conn.Exec(sql)
}

func (c *ClickHouse) Exec(sql string) {
	c.Conn.Exec(sql)
}

func (c *ClickHouse) QueryRow() {
	rows, err := c.Conn.Query("SELECT * FROM t_stock")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	stocks := make([]Stock, 0)
	for rows.Next() {
		var stock Stock
		rows.Scan(&stock.ID, &stock.Sku, &stock.Total, &stock.CreateTime)
		stocks = append(stocks, stock)
	}
	fmt.Println(stocks)
}
