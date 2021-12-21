package main

import (
	"log"

	"github.com/dgraph-io/badger/v3"
)

type Loader struct {
	db *badger.DB
}

func NewLoader() *Loader {
	opt := badger.DefaultOptions("/tmp/badger")
	opt.InMemory = true
	db, err := badger.Open(opt)
	if err != nil {
		log.Fatal(err)
	}
	return &Loader{db: db}
}

func (l *Loader) Close() {
	l.db.Close()
}

func (l *Loader) Get(key string) []byte {
	var valCopy []byte
	l.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			panic(err)
		}
		item.Value(func(val []byte) error {
			valCopy = append(valCopy, val...)
			return nil
		})
		return nil
	})
	return valCopy
}

func main() {
	loader := NewLoader()
	defer loader.Close()
}
