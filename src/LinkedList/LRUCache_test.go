package LinkedList

import (
	"testing"
)

func TestLRUCache(t *testing.T) {
	cache := NewLRUCache( 2 /* 缓存容量 */ );

	cache.Put(1, 1)
	t.Log(&cache)
	cache.Put(2, 2)
	t.Log(&cache)
	cache.Get(1)       // 返回  1
	t.Log(&cache)
	cache.Put(3, 3)    // 该操作会使得密钥 2 作废
	t.Log(&cache)
	cache.Get(2)       // 返回 -1 (未找到)
	t.Log(&cache)
	cache.Put(4, 4)    // 该操作会使得密钥 1 作废
	t.Log(&cache)
	cache.Get(1)       // 返回 -1 (未找到)
	t.Log(&cache)
	cache.Get(3)       // 返回  3
	t.Log(&cache)
	cache.Get(4)       // 返回  4
	t.Log(&cache)
}
