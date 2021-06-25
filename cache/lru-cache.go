package cache

import (
	"container/list"
	"sync"
)

type Cache interface {
	Set()
	Get()
	Del()
	Status()
}

type LRUCache struct {
	sm *sync.Map
	lData *list.List
	maxSize int
	misses, hits int64
}

type Element struct {
	k interface{}
	v interface{}
}


// 需要考虑值类型大小
// 决定key是否换为摘要
func NewLRUCache(size int) *LRUCache {
	return &LRUCache{
		maxSize: size,
		lData: list.New(),
	}
}

func (c *LRUCache)Get() {

}

func (c *LRUCache)Set()  {
	if
}

func (c *LRUCache) Status()  {

}

func (c *LRUCache) Del() {

}

