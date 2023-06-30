package main

import (
	"fmt"
)

func addElem(s []int) {
	s[0] = 100
	fmt.Printf("s: %v\n", s)
	// append触发扩容时，原来的指针指向的地址会发生变化，之后再对数组值进行改变，原切片将不受影响
	s = append(s, 4)
	fmt.Printf("s: %v\n", s)
	s[1] = 200
	fmt.Printf("s: %v\n", s)
}
