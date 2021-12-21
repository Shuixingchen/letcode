package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Loader struct {
	db  *sql.DB
	dsn string
}
type Cell struct {
	address    string
	capability int64
}

func NewLoader() Loader {
	dsn := "root:123456@tcp(127.0.0.1:3306)/forge?charset=utf8"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	return Loader{db: db, dsn: dsn}
}

func (l *Loader) QueryRow() {
	var cell Cell
	row := l.db.QueryRow("select address,capacity from address_unspent_tx_output where tx_hash=? and output_index=?", "aa", 0)
	err := row.Scan(&cell.address, &cell.capability)
	if err != nil {
		fmt.Println(err)
	}
	return
}
func (l *Loader) Query() []Cell {
	cells := make([]Cell, 0)
	rows, err := l.db.Query("select address,capacity from address_unspent_tx_output where tx_hash=? and output_index=?", "aa", 0)
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var cell Cell
		err := rows.Scan(&cell.address, &cell.capability)
		if err != nil {
			fmt.Println(err)
		}
		cells = append(cells, cell)
	}
	return cells
}

func main() {
	loader := NewLoader()
	fmt.Println(loader)
}
