package minidb

import "os"

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
	return &MiniDB{LogFile: logFile}, nil
}
