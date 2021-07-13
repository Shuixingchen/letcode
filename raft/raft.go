package main

import "fmt"

type EventLog struct {
	TimeStamp    int `json:"timestamp,omitempty"`
	ContractAddr string    `json:"address"`
	Name string `json:"name"`
}
func main(){
	var list []*EventLog
	newList := list[:6]
	fmt.Println(newList)
}