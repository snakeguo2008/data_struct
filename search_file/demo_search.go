package search_file

import (
	"fmt"
)

func SearchDemo() {
	array := []int{1, 3, 5, 6, 7, 8, 9, 10}

	index := Binary_search(array, 5)
	fmt.Println(index)
	//递归查找
	index_cycle := Binary_search_cycle(0, len(array)-1, array, 5)
	fmt.Println(index_cycle)

	//存在重复的元素
	array_1 := []int{1, 3, 4, 5, 6, 8, 8, 8, 11, 18}
	index_1 := Binary_search_first(array_1, 8)
	fmt.Println(index_1)

	array_2 := []int{1, 1, 3, 4, 5, 6, 8, 8, 8, 11, 18}
	index_2 := Binary_search_first(array_2, 1)
	fmt.Println(index_2)

	array_3 := []int{1, 1, 3, 4, 5, 6, 8, 8, 8, 11, 18, 18, 18}
	index_3 := Binary_search_end(array_3, 18)
	fmt.Println(index_3)

	//查找第一个大于等于目标元素的索引
	array_4 := []int{1, 3, 8, 8, 8, 8, 11, 18}
	index_4 := Binary_search_be(array_4, 8)
	fmt.Println(index_4)
}
