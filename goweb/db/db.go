package db

type DBInterface interface {
	Set(key, val []byte) error
	Get(key []byte) ([]byte, error)
	Del(key []byte) error
	Open(path string, sync bool) error
	GetAll() (int, error)
	Close() error
}
