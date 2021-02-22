package bitcoin

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

func GetSHA256HashCode(data interface{}) string{
	message,err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	hash := sha256.New()
	hash.Write(message)
	//计算哈希值
	bytes := hash.Sum(nil)
	//将字符串编码为16进制格式,返回字符串
	hashCode := hex.EncodeToString(bytes)
	//返回哈希值
	return hashCode
}

