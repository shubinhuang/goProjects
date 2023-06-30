package main

import "sync"

// 基于普通map实现一个线程安全的map，实现Add(k,v)和Delete(k)方法。
type mySafeMap struct {
	sync.RWMutex
	m map[int]int
}

func newSafeMap(len int) *mySafeMap {
	return &mySafeMap{
		m: make(map[int]int),
	}
}

func (mymap *mySafeMap) Add(k int, v int) {
	mymap.Lock()
	mymap.m[k] = v
	mymap.Unlock()
}

func (mymap *mySafeMap) Delete(k int) {
	mymap.Lock()
	delete(mymap.m, k)
	mymap.Unlock()
}

func (mymap *mySafeMap) Get(k int) (int, bool) {
	mymap.RLock()
	v, ok := mymap.m[k]
	mymap.RUnlock()
	return v, ok
}

func (mymap *mySafeMap) Set(k int, v int) {
	mymap.Lock()
	mymap.m[k] = v
	mymap.Unlock()
}
