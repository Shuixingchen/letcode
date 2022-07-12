package main

import (
	"database/sql"
	"fmt"
	"strconv"

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
	dsn := "root:123456@tcp(127.0.0.1:3306)/eth_parser?charset=utf8&maxAllowedPacket=17108864"
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

//通过预处理执行sql DDl语句
func (l *Loader) Prepare() {
	sqlInsert := "INSERT Into persons(PersonID, LastName, FirstName, Address, City)values(?, ?, ?, ?, ?)"
	stmt, err := l.db.Prepare(sqlInsert) //编译sql，并且返回一个stmt
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(1, "aa", "aa", "aa", "aa") //给这个stmt传参并且执行
	_, err = stmt.Exec(2, "bb", "bb", "bb", "bb") //给同一个stmt传参并且执行
	if err != nil {
		fmt.Println(err)
	}
}

//通过预处理执行query sql
func (l *Loader) PrepareQ() {
	sqlQuery := "select * from block where height = ? "
	stmt, err := l.db.Prepare(sqlQuery)
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()
	_, err = stmt.Query(1) //对同一个stmt，可以通过传不同的参数，执行不同的sql。
	_, err = stmt.Query(2)
	if err != nil {
		fmt.Println(err)
	}
}

//事务处理
func (l *Loader) Tx() {
	tx, err := l.db.Begin()
	sql1 := "INSERT Into persons(PersonID, LastName, FirstName, Address, City)values(?, ?, ?, ?, ?)"
	sql2 := "INSERT Into db_info(name, value)values(?, ?)"
	_, err = tx.Exec(sql1, "cc", "cc", "cc", "cc", "cc") //
	_, err = tx.Exec(sql2, "name1", "value1")
	if err != nil {
		tx.Rollback()
		fmt.Println(err)
	}
	tx.Commit()
}

//使用预处理语句处理事务
func (l *Loader) Txs() {
	sql2 := "INSERT Into db_info(name, value)values(?, ?)"
	tx, err := l.db.Begin()         //开启一个事务
	stmt, err := l.db.Prepare(sql2) //开启一个stmt,编译sql
	for i := 0; i < 100; i++ {
		_, err = tx.Stmt(stmt).Exec(strconv.Itoa(i), strconv.Itoa(i)) //在事务中执行已存在的stmt语句
		if err != nil {
			tx.Rollback()
		}
	}
	tx.Commit()
}
func (l *Loader) Handler() {
	sql1 := "select id from block_tmp where id not in (select MAX(id) from block_tmp group by block_number)"
	rows, err := l.db.Query(sql1)
	if err != nil {
		fmt.Println(err)
	}
	ids := make([]int, 0)
	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		if err != nil {
			fmt.Println(err)
		}
		ids = append(ids, id)
	}
	var str string
	for _, i := range ids {
		str += strconv.Itoa(i) + ","
	}
	fmt.Println(str)
}

func (l *Loader) QueryTx() {
	str := "select * from transactions where block_number in (0)"
	row := l.db.QueryRow(str)

	fmt.Println(row)
}

func main() {
	loader := NewLoader()
	loader.QueryTx()
}
