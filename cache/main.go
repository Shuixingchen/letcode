package main

import (
	"fmt"
	"letcode/cache/redis"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/patrickmn/go-cache"
)

var zkey = "lala"

func main() {
	c := cache.New(5*time.Minute, 10*time.Minute)

	// Set the value of the key "foo" to "bar", with the default expiration time
	c.Set("foo", "bar", cache.DefaultExpiration)

	// Set the value of the key "baz" to 42, with no expiration time
	// (the item won't be removed until it is re-set, or removed using
	// c.Delete("baz")
	c.Set("baz", 42, cache.NoExpiration)

	// Since Go is statically typed, and cache values can be anything, type
	// assertion is needed when values are being passed to functions that don't
	// take arbitrary types, (i.e. interface{}). The simplest way to do this for
	// values which will only be used once--e.g. for passing to another

	type MyStruct struct {
		ID int
	}
	var m MyStruct
	m.ID = 1
	c.Set("foo", &m, cache.DefaultExpiration)
	if x, found := c.Get("foo"); found {
		foo := x.(*MyStruct)
		fmt.Println(foo)
	}
}

func InitZSet(rc *redis.Client, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		t := time.Now().Unix()
		_, err := rc.ZAdd(zkey, float64(t), strconv.FormatInt(t, 10))
		if err != nil {
			log.Fatal(err)
			return
		}
		time.Sleep(1 * time.Second)
	}
}
