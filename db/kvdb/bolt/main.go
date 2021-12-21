package main

import (
	"fmt"
	"log"
	"time"

	"github.com/boltdb/bolt"
)

/*
 */

type KDb struct {
	db     *bolt.DB
	bucket string
}

type Simple struct {
	db *bolt.DB
}

func NewKDb() *KDb {
	options := &bolt.Options{Timeout: 10 * time.Second}
	db, err := bolt.Open("/home/csx/letcode/db/kvdb/bolt/data/my.db", 0600, options)
	if err != nil {
		log.Fatal(err)
	}
	return &KDb{db: db}
}
func NewSimpleDb() *Simple {
	options := &bolt.Options{Timeout: 10 * time.Second}
	db, err := bolt.Open("/home/csx/letcode/db/kvdb/bolt/data/simple.db", 0600, options)
	if err != nil {
		log.Fatal(err)
	}
	return &Simple{db: db}
}

func (k *KDb) Init(bucket string) {
	k.db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
	k.bucket = bucket
}
func (k *KDb) Add(key string, value []byte) error {
	return k.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(k.bucket))
		err := b.Put([]byte(key), value)
		return err
	})
}
func (k *KDb) Get(key string) []byte {
	var res []byte
	k.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(k.bucket))
		res = b.Get([]byte(key))
		return nil
	})
	return res
}
func (k *KDb) Delete(key string) error {
	return k.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(k.bucket))
		return b.Delete([]byte(key))
	})
}

func (k *KDb) Close() {
	k.db.Close()
}

func main() {
	k := NewKDb()
	k.Init("user")
	k.Add("key", []byte("value"))
	k.Add("key", []byte("value1212"))
	res := k.Get("key")
	k.Delete("key")
	fmt.Println(string(res))
}
