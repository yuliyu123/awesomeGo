package lru

import "container/list"

type Value interface {
	Len() int
}

type entry struct {
	key string
	value Value
}

type Cache struct {
	maxBytes int64
	nBytes int64
	ll *list.List
	cache map[string]*list.Element
	onEvicted func(key string, value Value)
}

// func func_name return_value
// New is the Constructor of Cache
func New(maxBytes int64, onEvicted func(key string, value Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		//nBytes:    0,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		onEvicted: onEvicted,
	}
}

// Get look ups a key's value
func (c *Cache) Get(key string) (value Value, ok bool) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele) // move to front, latest hot data
		kv := ele.Value.(*entry)
		return kv.value, true
	}
	return
}

// delete the oldest element
func (c* Cache) RemoveTheOldest() {
	ele := c.ll.Back()
	if ele != nil {
		kv := ele.Value.(*entry)
		c.ll.Remove(ele)
		delete(c.cache, kv.key)
		c.nBytes -= int64(len(kv.key)) + int64(kv.value.Len())
		if c.onEvicted != nil {
			c.onEvicted(kv.key, kv.value)
		}
	}
}

// add a new entry
func (c* Cache) AddToCache(key string, value Value) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		c.nBytes += int64(value.Len()) - int64(kv.value.Len())
	} else {
		ele := c.ll.PushFront(&entry{key, value})
		c.cache[key] = ele
		c.nBytes += int64(len(key)) + int64(value.Len())
	}
		for c.maxBytes != 0 && c.maxBytes < c.nBytes { // when maxBytes < nBytes(used), remove the oldest
		c.RemoveTheOldest()
	}
}

func (c *Cache) len() int {
	return c.ll.Len()
}

//func main() {
//
//}