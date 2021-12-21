package main

import (
	"encoding/json"
	"fmt"
)

type Stu struct {
	Name string `json:"name"` //tag标签备注json，对应的key
	Key  string
}

func main() {
	m := make(map[string]Stu)
	m["ss"] = Stu{"aa", "ss"}
	m["dd"] = Stu{"aa", "dd"}
	res, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
	}
	Decode(string(res))
}

func Decode(data string) {
	m := make(map[string]Stu)
	if err := json.Unmarshal([]byte(data), &m); err != nil {
		fmt.Println(err)
	}
	fmt.Println(m)
}
