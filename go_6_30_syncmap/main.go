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
	}

	wg.Wait()
}

func testMySafeMap() {
	num := 100

	mymap := newSafeMap(num)
	wg := sync.WaitGroup{} // 让代码等待一组 goroutine 的结束
	wg.Add(num)

	for i := 0; i < num; i++ {
		go func(num int) {
			mymap.Add(num, num)
			wg.Done()
		}(i)
	}

	wg.Wait()
}

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
	// testSlice()
	testUnSafeMap()
	testMySafeMap()
	// test_1()
}
