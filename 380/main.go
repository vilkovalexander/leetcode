package main

import (
	"fmt"
	"math/rand"
)

type RandomizedSet struct {
	hashmap map[int]int
	list    []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{make(map[int]int, 0), make([]int, 0)}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.hashmap[val]; ok {
		return false
	}
	this.hashmap[val] = len(this.list)
	this.list = append(this.list, val)
	return true

}

func (this *RandomizedSet) Remove(val int) bool {
	if index, ok := this.hashmap[val]; ok {
		if index != len(this.list)-1 {
			this.hashmap[this.list[len(this.list)-1]] = index
			this.list[index], this.list[len(this.list)-1] = this.list[len(this.list)-1], this.list[index]
		}
		this.list = this.list[:len(this.list)-1]
		delete(this.hashmap, val)
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	return this.list[rand.Intn(len(this.list))]
}

func main() {
	r := Constructor()
	fmt.Println(r.Insert(1))
	fmt.Println(r.Remove(2))
	fmt.Println(r.Insert(2))
	fmt.Println(r.GetRandom())
	fmt.Println(r.Remove(1))
	fmt.Println(r.Insert(2))
	fmt.Println(r.GetRandom())
}
