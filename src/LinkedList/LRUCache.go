package LinkedList

import (
	"container/list"
	"fmt"
)

type LRUCache struct {
	capacity int
	l *list.List
	m map[interface{}]*list.Element
	length int
}

type entry struct {
	key interface{}
	value interface{}
}


func NewLRUCache(capacity int) LRUCache {
	return LRUCache{capacity,list.New(),make(map[interface{}]*list.Element, capacity), 0}
}


func (this *LRUCache) Get(key int) int {
	if v, ok := this.m[key]; ok {
		this.l.MoveToFront(v)
		return v.Value.(*entry).value.(int)
	}
	return -1
}


func (this *LRUCache) Put(key int, value int)  {
	if v, ok := this.m[key]; ok {
		this.l.MoveToFront(v)
		v.Value.(*entry).value = value
		return
	}

	ent := &entry{key, value}
	lent := this.l.PushFront(ent)
	this.m[key] = lent

	if this.length == this.capacity {
		this.removeOldest()
	} else {
		this.length++
	}

	return
}

func (this *LRUCache) removeOldest() {
	lentry := this.l.Back()
	if nil != lentry {
		key := lentry.Value.(*entry).key
		delete(this.m, key)
		this.l.Remove(lentry)
	}
}

func (this *LRUCache) String() string {
	var ret string
	cur := this.l.Front()
	for nil != cur {
		ret += fmt.Sprintf("%+v", cur.Value)
		if nil != cur.Next() {
			ret += "->"
		}
		cur = cur.Next()
	}
	return ret
}