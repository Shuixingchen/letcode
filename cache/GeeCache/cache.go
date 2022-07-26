package geecache

import (
	"container/list"
	"errors"
)

/*
LRU 算法内存缓存
map: 存储键与链表节点指针的映射
list: Go contain的标准双向链表实现，节点存储key,value
每次访问一个key,把这个节点移动到队尾
每次插入一个key,在队尾新增一个节点
如果达到最大容量，删除队头的节点
*/

type Cache struct {
	MaxSize int
	ll      *list.List
	cache   map[string]*list.Element
	nbyte   int
}

type entry struct {
	key   string
	value Value
}
type Value interface {
	Len() int
}

func NewCache() *Cache {
	return &Cache{
		MaxSize: 10 * 1024 * 1024,
		ll:      list.New(),
		cache:   make(map[string]*list.Element),
	}
}

func (c *Cache) Get(key string) (Value, error) {
	if e, ok := c.cache[key]; ok {
		c.ll.MoveToFront(e)
		kv := e.Value.(*entry)
		return kv.value, nil
	}
	return nil, errors.New("key not exists")
}

func (c *Cache) Set(key string, value Value) {
	if e, ok := c.cache[key]; ok {
		c.ll.MoveToFront(e)
		kv := e.Value.(*entry)
		c.nbyte += value.Len() - kv.value.Len()
	} else {
		c.ll.PushFront(entry{key: key, value: value})
		c.nbyte += len(key) + value.Len()
	}
	if c.MaxSize != 0 && c.MaxSize < c.nbyte {
		c.MoveOld()
	}
	return
}

func (c *Cache) MoveOld() {
	last := c.ll.Back()
	if last != nil {
		c.ll.Remove(last)
		kv := last.Value.(*entry)
		delete(c.cache, kv.key)
		c.nbyte -= len(kv.key) + kv.value.Len()
	}
}
