package minidb

import "os"

const LSMLogName = "minidb.db"

type LogFile struct {
	File   *os.File
	Offset int64
}

func OpenLogFile(path string) (f *LogFile, err error) {
	f.File, err = os.OpenFile(path+LSMLogName, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return nil, err
	}
	stat, err := os.Stat(path + LSMLogName)
	if err != nil {
		return nil, err
	}
	f.Offset = stat.Size()
	return
}

func (f *LogFile) Read(offset int64) (e *Entry, err error) {
	buf := make([]byte, entryHeaderSize)
	if _, err = f.File.ReadAt(buf, offset); err != nil {
		return
	}
	if e, err = DecodeHeader(buf); err != nil {
		return
	}
	offset += entryHeaderSize
	if e.KeySize > 0 {
		key := make([]byte, e.KeySize)
		if _, err = f.File.ReadAt(key, offset); err != nil {
			return
		}
		e.Key = key
	}
	offset += int64(e.KeySize)
	if e.ValueSize > 0 {
		value := make([]byte, e.ValueSize)
		if _, err = f.File.ReadAt(value, offset); err != nil {
			return
		}
		e.Value = value
	}
	return
}

func (f *LogFile) Write(e *Entry) error {
	buf := e.Encode()
	if _, err := f.File.WriteAt(buf, f.Offset); err != nil {
		return err
	}
	f.Offset += e.GetSize()
	return nil
}
