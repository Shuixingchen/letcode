package redis

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"sync"
	"time"
)

const (
	expiry=1*time.Minute
)
// Client is a struct implements the reader interface and has actual redis client embeded.
type Client struct {
	rc *redis.Client
	m  *sync.Mutex
}

func NewRedisClient() (*Client, error){
	return New("localhost:6379","",0)
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

func (rc *Client)AddDataToCache( key string, data interface{}) error{
	datastr, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if _, err := rc.Set(key, string(datastr), expiry); err != nil {
		return err
	}
	return nil
}
func (rc *Client)GetDatafromCache(key string, data interface{}) error{
	value, err := rc.Get(key)
	if err != nil {
		return  err
	}
	err = json.Unmarshal([]byte(value),data)
	if err !=nil{
		return err
	}
	return nil
}
func (rc *Client)AddDataToCacheNew( key string, data interface{}) error{
	datastr,err := Serialize(data)
	if err != nil {
		return err
	}
	if _, err := rc.Set(key, string(datastr), expiry); err != nil {
		return err
	}
	return nil
}
func (rc *Client)GetDatafromCacheNew(key string, data interface{}) error{
	value, err := rc.Get(key)
	if err != nil {
		return  err
	}
	err = Deserialize([]byte(value), data)
	if err !=nil{
		return err
	}
	return nil
}
