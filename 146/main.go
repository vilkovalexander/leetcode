package main

import "container/list"

type LRUCache struct {
	mainList *list.List
	dict     map[int]*list.Element
	n        int
}

type pair struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{n: capacity, dict: make(map[int]*list.Element, 0), mainList: list.New()}
}

func (this *LRUCache) Get(key int) int {
	if val, ok := this.dict[key]; ok {
		this.mainList.MoveToFront(val)
		return val.Value.(pair).value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if val, ok := this.dict[key]; ok {
		val.Value = pair{key: key, value: value}
		this.mainList.MoveToFront(val)
		return
	}
	if this.mainList.Len() < this.n {
		el := this.mainList.PushFront(pair{key: key, value: value})
		this.dict[key] = el
		return
	}
	delete(this.dict, this.mainList.Back().Value.(pair).key)
	this.mainList.Remove(this.mainList.Back())
	el := this.mainList.PushFront(pair{key, value})
	this.dict[key] = el
	return
}
