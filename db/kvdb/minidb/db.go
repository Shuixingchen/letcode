package minidb

import (
	"errors"
	"os"
)

// 基于LSM树实现kv数据库

type MiniDB struct {
	LogFile *LogFile
	Indexs  map[string]int64
	path    string
}

type Options struct {
}

func Open(path string) (*MiniDB, error) {
	// 如果目录不存在，就新建一个
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err = os.MkdirAll(path, os.ModePerm); err != nil {
			return nil, err
		}
	}
	// 加载数据文件
	logFile, err := OpenLogFile(path)
	if err != nil {
		return nil, err
	}
	db := &MiniDB{
		LogFile: logFile,
		Indexs:  make(map[string]int64),
		path:    path,
	}
	db.LoadIndexFromFile()
	return db, nil
}

func (db *MiniDB) LoadIndexFromFile() {

}

func (db *MiniDB) Set(key, value []byte) error {
	offset := db.LogFile.Offset
	// 封装entry
	entry := NewEntry(key, value, SET)
	// 写入log文件
	err := db.LogFile.Write(entry)
	if err != nil {
		return err
	}
	// 写到内存索引
	db.Indexs[string(key)] = offset
	return nil
}

func (db *MiniDB) Get(key []byte) (value []byte, err error) {
	offset, ok := db.Indexs[string(key)]
	if !ok {
		return []byte{}, errors.New("key not exist")
	}
	e, err := db.LogFile.Read(offset)
	if err != nil {
		return []byte{}, err
	}
	return e.Value, nil
}

func (db *MiniDB) Merge() {

}
