package main

import (
	"fmt"
	"letcode/cache/redis"
	"log"
	"strconv"
	"sync"
	"time"
)

var zkey = "lala"

func main() {
	rc, err := redis.New("127.0.0.1:6379", "123456", 0)
	if err != nil {
		log.Fatal(err)
	}
	var wg sync.WaitGroup
	wg.Add(2)
	now := time.Now().Unix()
	var delStart int64
	delStart = 1641716417
	go InitZSet(rc, &wg)

	time.Sleep(5 * time.Second)
	res4, err := rc.ZRevRange(zkey, 0, 5)
	res5, err := rc.ZCount(zkey, strconv.FormatInt(now, 10), strconv.FormatInt(now+5, 10))
	res6, err := rc.ZRemRangeByScore(zkey, strconv.FormatInt(delStart, 10), strconv.FormatInt(now-60, 10))

	fmt.Println(res4, res5, res6)
	wg.Wait()
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
