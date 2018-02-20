package set

import (
	"fmt"
	"sync"
)

var pool = sync.Pool{}

type Set struct {
	items map[interface{}]struct{}
	lock  sync.RWMutex
}

func (set *Set) Add(items ...interface{}) {
	set.lock.Lock()
	defer set.lock.Unlock()

	for _, v := range items {
		set.items[v] = struct{}{}
	}
}

func (set *Set) Remove(items ...interface{}) {
	set.lock.Lock()
	defer set.lock.Unlock()

	for _, v := range items {
		delete(set.items, v)
	}
}

func New() *Set {
	s := Set{}
	s.items = make(map[interface{}]struct{})
	return &s
}

func (set Set) String() string {
	return fmt.Sprintf("%v", set.items)
}
