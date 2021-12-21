package main

import "container/list"

/*
LRU策略缓存，维护一个队列，如果缓存被访问，把它移动到队尾，那么队首就是最近最少访问的缓存。
*/

type Cache struct {
	maxSize   int64
	nSize     int64      //已经使用的内存
	ll        *list.List //标准库的双向队列
	cache     map[string]*list.Element
	OnEvicted func(key string, value Value) //某个缓存被移除后的回调函数
}

type entry struct {
	key   string
	value Value
}
type Value interface {
	Len() int64
}

func New(maxSize int64, onEvicted func(key string, value Value)) *Cache {
	return &Cache{
		maxSize:   maxSize,
		nSize:     0,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}
func (c *Cache) RemoveOld() {
	ele := c.ll.Back()
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key)
		c.nSize -= int64(len(kv.key)) + int64(kv.value.Len())
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}
}

func (c *Cache) Add(key string, value Value) {

}
