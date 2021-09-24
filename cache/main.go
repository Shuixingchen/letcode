package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	var m sync.Map
	for i := 0; i < 10; i++ {
		m.Store(i, strconv.Itoa(i))
	}

	v, ok := m.Load(1)
	if ok == false {
		v = "default"
	}

	fmt.Println("v:", v.(string))
}
