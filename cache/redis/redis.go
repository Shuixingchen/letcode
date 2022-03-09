package redis

import (
	"sync"
	"time"

	"github.com/go-redis/redis"
)

// Client is a struct implements the reader interface and has actual redis client embeded.
type Client struct {
	rc *redis.Client
	m  *sync.Mutex
}

// New returns a new pointer to Client object given the connection details.
func New(addr string, pwd string, db int) (*Client, error) {
	rc := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd,
		DB:       db,
	})
	_, err := rc.Ping().Result()
	if err != nil {
		return nil, err
	}
	var mutex = &sync.Mutex{}

	return &Client{rc, mutex}, nil
}

// Get returns a key's corresponding value from redis and error if any.
func (c *Client) Get(key string) (string, error) {
	return c.rc.Get(key).Result()
}

// Set adds a new key/value pair along with expiration to redis server.
func (c *Client) Set(key string, value string, expr time.Duration) (string, error) {
	return c.rc.Set(key, value, expr).Result()
}

// LPush Insert all the specified values at the head of the list stored at key.
func (c *Client) LPush(key string, values ...interface{}) (int64, error) {
	return c.rc.LPush(key, values...).Result()
}

func (c *Client) LPop(key string) (string, error) {
	return c.rc.LPop(key).Result()
}

// Exists checks if redis exists a key or not,
// 1 means exist.
// 0 means not exist.
func (c *Client) Exists(key string) (int64, error) {
	return c.rc.Exists(key).Result()
}

// Sort returns or stores the elements contained in the list, set or sorted set at key.
func (c *Client) Sort(key string, sort *redis.Sort) ([]string, error) {
	return c.rc.Sort(key, sort).Result()
}

// LRem removes a value in given list.
func (c *Client) LRem(list string, count int64, value string) (int64, error) {
	return c.rc.LRem(list, count, value).Result()
}

// Keys returns keys matched a given pattern from redis.
func (c *Client) Keys(pattern string) ([]string, error) {
	return c.rc.Keys(pattern).Result()

}

// Del deletes a key from the redis server.
func (c *Client) Del(keys ...string) (int64, error) {
	return c.rc.Del(keys...).Result()
}

// Close shutdown the connection with underlying redis server.
func (c *Client) Close() error {
	return c.rc.Close()
}

// LLen returns given list len in redis server.
func (c *Client) LLen(key string) (int64, error) {
	return c.rc.LLen(key).Result()
}

// Lock calls the mutex lock method to protect multi-thread read/write.
func (c *Client) Lock() {
	c.m.Lock()
}

// Unlock calls the mutex unlock method to protect multi-thread read/write.
func (c *Client) Unlock() {
	c.m.Unlock()
}

// FlushAll will flush all the existing keys stored in redis.
func (c *Client) FlushAll() {
	c.rc.FlushAll()
}

// ZAdd adds a new member with given score to specified sorted list.
func (c *Client) ZAdd(key string, score float64, item string) (int64, error) {
	member := redis.Z{Score: score, Member: item}
	return c.rc.ZAdd(key, member).Result()
}

// 删除有序集合元素
func (c *Client) ZRem(key string, item string) (int64, error) {
	return c.rc.ZRem(key, item).Result()
}

// 按照score倒序获取列表
func (c *Client) ZRevRange(key string, start, stop int64) ([]string, error) {
	return c.rc.ZRevRange(key, start, stop).Result()
}

// 获取指定范围的个数
func (c *Client) ZCount(key, min, max string) (int64, error) {
	return c.rc.ZCount(key, min, max).Result()
}

// 移除指定score区间的所有成员
func (c *Client) ZRemRangeByScore(key, min, max string) (int64, error) {
	return c.rc.ZRemRangeByScore(key, min, max).Result()
}
