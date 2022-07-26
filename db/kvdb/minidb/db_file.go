package minidb

import "os"

const LSMLogName = "minidb.db"

type LogFile struct {
	File   *os.File
	Offset int
}

type Entry struct {
	Key   string
	Value []byte
	Type  int
}

func OpenLogFile(path string) (f *LogFile, err error) {
	f.File, err = os.OpenFile(path+LSMLogName, os.O_CREATE, 0666)
	return
}
