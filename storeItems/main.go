package main

import (
	"container/list"
	"fmt"
)

const CacheSize = 20


type entry struct {
	key   string
	value string
}


type LRUCache struct {
	capacity int
	cache    map[string]*list.Element
	list     *list.List 
}


func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[string]*list.Element),
		list:     list.New(),
	}
}


func (l *LRUCache) Get(key string) (string, bool) {
	if element, exists := l.cache[key]; exists {
		l.list.MoveToFront(element) 
		return element.Value.(entry).value, true
	}
	return "", false
}


func (l *LRUCache) Put(key, value string) {
	if element, exists := l.cache[key]; exists {
		l.list.MoveToFront(element)
		element.Value = entry{key, value}
		return
	}

	
	if l.list.Len() == l.capacity {
		lastElement := l.list.Back()
		if lastElement != nil {
			evictedEntry := lastElement.Value.(entry)
			delete(l.cache, evictedEntry.key)
			l.list.Remove(lastElement)
			fmt.Println("Evicted:", evictedEntry.key)
		}
	}

	
	newElement := l.list.PushFront(entry{key, value})
	l.cache[key] = newElement
	fmt.Println("Added:", key)
}

func main() {
	cache := NewLRUCache(CacheSize)

	
	for i := 1; i <= 20; i++ {
		key := fmt.Sprintf("key%d", i)
		value := fmt.Sprintf("value%d", i)
		cache.Put(key, value)
	}

	
	cache.Get("key5")
	cache.Get("key10")

	
	cache.Put("key21", "value21")


	if val, ok := cache.Get("key5"); ok {
		fmt.Println("Get key5:", val)
	} else {
		fmt.Println("key5 not found")
	}

	if val, ok := cache.Get("key1"); ok {
		fmt.Println("Get key1:", val)
	} else {
		fmt.Println("key1 evicted (Least Recently Used)")
	}
}
