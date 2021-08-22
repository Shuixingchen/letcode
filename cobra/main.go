package main

import "fmt"

type base struct {
	Name string
}

func (b *base) Add() {
	b.Name = b.Name + "ddd"
}

type Newst struct {
	base
}

func (n *Newst) Add() {
	n.Name = "bbb" + "ccc"
}

func main() {
	var str Newst
	str.Name = "aa"
	str.Add()

	fmt.Println(str.Name)
}
