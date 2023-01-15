package go_algorithms

import (
	"errors"
)

// @program:     go_algorithms
// @file:        selectSorter_test.go
// @author:      bowen
// @create:      2022-11-12 15:24
// @description: 选择排序

func selectionSort(arr []int) error {
	// 空数组和1个元素的不需要排序
	if arr == nil || len(arr) < 2 {
		return errors.New("空数组和1个元素的数组不需要排序")
	}

	// n>=2
	n := len(arr)
	for i := 0; i < n-1; i++ {

		arr[i] = min
	}
	return nil
}
