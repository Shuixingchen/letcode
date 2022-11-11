package models

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"math/big"
	"math/rand"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

type TokenHolder struct {
	TokenAddr    string
	HolderAddr   string
	TokenBalance *big.Int
}

func CreateTokenHolder() string {
	sql := `CREATE TABLE t_token_holder
	(
		"token_addr" String,
		"holder_addr" String,
		"token_balance" Int128,
	)ENGINE = ReplacingMergeTree 
	ORDER BY (holder_addr,token_addr);`
	return sql
}

// insert into table () values('str',100), 如果字段是string,一定要用单引号
func InsertTokenHolder(tokenHolders []*TokenHolder) string {
	insertPrefix := "Insert into t_token_holder(token_addr, holder_addr, token_balance) VALUES "
	var builder strings.Builder
	for _, tokenHolder := range tokenHolders {
		builder.WriteString(fmt.Sprintf("('%s', '%s', %d)", tokenHolder.HolderAddr, tokenHolder.TokenAddr, tokenHolder.TokenBalance))
	}
	return insertPrefix + builder.String()
}

func Query(conn *sql.DB, sql string) (tokenHolders []*TokenHolder) {
	rows, err := conn.Query(sql)
	if err != nil {
		log.WithField("sql", sql).Error(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var t TokenHolder
		rows.Scan(&t.HolderAddr, &t.TokenAddr, &t.TokenBalance)
		tokenHolders = append(tokenHolders, &t)
	}
	if err = rows.Err(); err != nil {
		log.WithField("sql", sql).Error(err)
		return
	}
	return tokenHolders
}

// 随机产生数据
func GenerateTokenHolders(num int) []*TokenHolder {
	res := make([]*TokenHolder, 0)
	for i := 0; i < num; i++ {
		r := rand.Intn(1000000000)
		tokenAddr := GetSHA256HashCode([]byte(strconv.Itoa(r)))
		for j := 0; j < 10; j++ {
			var t TokenHolder
			balance := big.NewInt(0)
			t.TokenAddr = tokenAddr
			t.HolderAddr = GetSHA256HashCode([]byte(strconv.Itoa(i + j*2)))
			t.TokenBalance, _ = balance.SetString("100000000000000000000", 10)
			res = append(res, &t)
		}
	}
	return res
}

/*
*
SHA256生成哈希值
*/
func GetSHA256HashCode(message []byte) string {
	//方法一：
	//创建一个基于SHA256算法的hash.Hash接口的对象
	hash := sha256.New()
	//输入数据
	hash.Write(message)
	//计算哈希值
	bytes := hash.Sum(nil)
	//将字符串编码为16进制格式,返回字符串
	hashCode := hex.EncodeToString(bytes)
	//返回哈希值
	return "0x" + hashCode[:40]
}
