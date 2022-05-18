package models

import (
	"context"
	"database/sql"
	"fmt"
	"math/big"
	"strings"
)

// TokenType represents the token type used in token model.
type TokenType = int

const (
	_ TokenType = iota
	// TokenERC20 represents the ERC20 Token type in ETH.
	TokenERC20
	// TokenERC721 represents the ERC721 Token type in ETH.
	TokenERC721
)

// Token - represents Token contract referred in ETH.
type Token struct {
	Addr            string    `json:"tokenAddress"`
	Name            string    `json:"name"`
	Symbol          string    `json:"sym"`
	Decimals        int64     `json:"decimals"`
	InitTotalSupply *big.Int  `json:"initTotalSupply"` // the maximum supply of a ERC20 token can be 2^256 - 1
	Type            TokenType `jsonn:"tokenType"`
}

// InsertionPrefix returns the token insertion sql query prefix string.
func (t Token) InsertionPrefix() string {
	return "INSERT INTO tokens (`token_address`, `name`, `symbol`, `decimals`, `init_total_supply`, `type`) VALUES"
}

// InsertionValue returns the Token insertion sql query value string.
func (t Token) InsertionValue() string {
	return fmt.Sprintf("('%s', '%s', '%s', '%d', '%s', '%d')",
		t.Addr,
		t.Name,
		t.Symbol,
		t.Decimals,
		t.InitTotalSupply.String(),
		t.Type,
	)
}

// InsertQuery returns the insertion string for given token.
func (t Token) InsertQuery() string {
	return t.InsertionPrefix() + t.InsertionValue()
}

// TokenBatchInsertion returns a batch insertion query given a token list.
func TokenBatchInsertion(list []*Token) string {
	if len(list) == 0 {
		return ""
	}
	prefix := list[0].InsertionPrefix()
	var inserts []string
	for _, item := range list {
		inserts = append(inserts, item.InsertionValue())
	}
	return prefix + strings.Join(inserts, ",") + `ON DUPLICATE KEY UPDATE 
	name=values(name), 
	symbol=values(symbol),
	decimals=values(decimals),
	init_total_supply=values(init_total_supply),
	type=values(type);`
}

func GetToken(ctx context.Context, query string, db *sql.DB) (*Token, error) {
	var token = new(Token)
	row := db.QueryRow(query)
	if err := row.Scan(
		&token.Addr,
		&token.Name,
		&token.Symbol,
		&token.Decimals,
		&token.InitTotalSupply,
		&token.Type,
	); err != nil {
		return nil, err
	}
	return token, nil
}
func GetTokenHash(ctx context.Context, query string, db *sql.DB) (string, error) {
	var hash string
	row := db.QueryRow(query)
	if err := row.Scan(
		&hash,
	); err != nil {
		return "", err
	}
	return hash, nil
}

// TokenSelectQueryPrefix returns the prefix for token select query.
const TokenSelectQueryPrefix = `select
token_address,
name,
symbol,
decimals,
init_total_supply,
type from tokens `
