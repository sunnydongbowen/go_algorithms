package go_algorithms

import (
	"fmt"
	"testing"
)

// @program:     go_algorithms
// @file:        bubblesorter.go
// @author:      bowen
// @create:      2022-11-26 20:03
// @description: 冒泡排序

func BublleSort(a []int) {
	var tmp int
	//
	for i := 0; i < len(a)-1; i++ {
		// 精华就在这,第二轮i=1的时候，其实少比较了一轮，因为已经比较好了。
		for j := 0; j < len(a)-1-i; j++ {
			// 这一段是可以自己写出来的
			if a[j] > a[j+1] {
				tmp = a[j]
				a[j] = a[j+1]
				a[j+1] = tmp
			}
		}
	}
}

func TestBublleSort(t *testing.T) {
	a := []int{5, 3, 2, 1, 6}
	BublleSort(a)
	fmt.Print(a)
}
