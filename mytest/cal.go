package main

func Mul(a, b int64) int64 {
	return a * b
}

func Add(a, b int64) int64 {
	return a + b
}

// db.go
type DB interface {
	Get(key string) (int, error)
}

func GetFromDB(db DB, key string) int {
	if value, err := db.Get(key); err == nil {
		return value
	}

	return -1
}
