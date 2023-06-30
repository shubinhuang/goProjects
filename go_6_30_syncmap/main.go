package main

import (
	"fmt"
	"sync"
)

func testSlice() {
	var s1 = make([]int, 0, 3)
	s1 = append(s1, 1, 2, 3)
	addElem(s1)
	fmt.Printf("s1: %v\n", s1)
}

// 普通map
func testUnSafeMap() {
	num := 100

	mymap := make(map[int]int, num)
	wg := sync.WaitGroup{} // 让代码等待一组 goroutine 的结束
	wg.Add(num)

	for i := 0; i < num; i++ {
		go func(num int) {
			mymap[num] = num
			wg.Done()
		}(i)
		go func(num int) {
			_, _ = mymap[num]
		}(i)
	}

	wg.Wait()
}

// 自己实现的线程安全的map
func testMySafeMap() {
	num := 100

	mymap := newSafeMap(num)
	wg := sync.WaitGroup{} // 让代码等待一组 goroutine 的结束
	wg.Add(num)            // 加num

	for i := 0; i < num; i++ {
		go func(num int) {
			mymap.Add(num, num)
			wg.Done() // 减1
		}(i)

		go func(num int) {
			_, _ = mymap.Get(num)
		}(i)
	}

	wg.Wait() // 等待为0返回
}

// sync.map
func testSyncMap() {
	num := 100

	var mymap sync.Map
	wg := sync.WaitGroup{} // 让代码等待一组 goroutine 的结束
	wg.Add(num)            // 加num

	for i := 0; i < num; i++ {
		go func(num int) {
			mymap.Store(num, num)
			wg.Done() // 减1
		}(i)

		go func(num int) {
			_, _ = mymap.Load(num)
		}(i)
	}

	wg.Wait() // 等待为0返回
}

// 测试基本功能
func test_1() {
	mymap := newSafeMap(0)
	mymap.Add(1, 1)
	mymap.Add(3, 3)
	mymap.Add(5, 5)
	fmt.Println(mymap)
	mymap.Set(1, 2)
	fmt.Println(mymap)
	mymap.Delete(3)
	fmt.Println(mymap)
	fmt.Println(mymap.Get(5))
}

func main() {
	testSlice()
	// testUnSafeMap()
	// testMySafeMap()
	// testSyncMap()
	// test_1()
}
