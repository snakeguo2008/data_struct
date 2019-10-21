package sort

/************
*** 1、简单选择排序
*** 2、每次选出最小的和对应的索引下标交换
***
 */

func SelectSort(array []uint32) {
	if len(array) == 0 {
		return
	}
	iLength := len(array)
	index := 0
	i, j := 0, 0

	for i = 0; i < iLength; i++ {

		index = i
		//在[i:] 数据中找到最小的
		for j = i + 1; j < iLength; j++ {
			if array[j] < array[index] {
				index = j
			}
		}
		//当前下标i对应的
		if index > i {
			array[i], array[index] = array[index], array[i]
		}
	}
}
