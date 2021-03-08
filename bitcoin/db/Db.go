package db


import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Db struct {
	db *gorm.DB
}
//
type Block struct {
	
}

func CreateDb() *Db{
	 db,err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	if err!= nil {
		fmt.Println(err)
	}
	 return &Db{db:db}
}

func (db *Db)Get(key []byte) {

}