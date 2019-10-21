package sort_practise

/***题目****
*** 设计一个算法,使得在尽可能少的时间内重排数组, 将所有负数的关键字
*** 放在取正数关键字的之前
***
 */

/***
*** 思路: 设定两个下标索引left=最左，right=最右,
*** 类似于快排一样交换两个正负数，以达到要求
***/

/***
1、时间复杂度：由于只是遍历了一遍，所以时间复杂度为O(n);空间复杂度为 O(1)
**************/

func PSwitch(array []int) {
	if len(array) <= 1 {
		return
	}

	i, j := 0, len(array)-1
	for i < j {
		for ; i < j && array[j] >= 0; j-- {
		}
		for ; i < j && array[i] <= 0; i++ {

		}
		if i < j {
			array[i], array[j] = array[j], array[i]
		}

	}

}

/***题目****
*** 设向量A[0,...,n-1]中存有n个互不相同的整数，每个关键字都在0~n-1。实现
*** 将向量A排序，结果输出到向量B[0,...,n-1]中
 */

/**思路*****
** 明摆着A中的数就是0~n-1的乱序排列，那么只用找该数和对应下标的数交换就行了
** 遍历次数1次，交换次数，最好的0次，最差的n-1次
** 当前位置交换中止的条件就是索引和数相同
***********/

//我实现不用辅助数组B就完成数组交换
func PSwitchNum(array []uint32) {

	iLen := len(array)
	for i := 0; i < iLen; i++ {

		for array[i] != uint32(i) {
			array[i], array[array[i]] = array[array[i]], array[i]
		}
	}

}

/***题目****
*** 计数排序。该算法对一个待排序的表(用数组A[]表示)进行排序，排序结果存储在另一个新表(B[])，表中关键字为
*** int类型。必须注意的是，标桩待排序的关键字互不相同，计数排序算法对表中的每个关键字，扫描待排序表一趟，
*** 统计表中有多少个关键字比该关键字小。假设对某个关键字，统计数字为c,那么该关键字在新的有序表中位置即为c
 */

/*****思路*****/
/*****
 *** 无非是根据找到该关键字在新表中的索引位置罢了
 ***
 *** 注意的是每次都要扫描全表
 *** 这样比较次数是n^2. 时间复杂度是O(n^2),空间复杂度是n
 *** 简单选择排序比较次数是n(n-1)/2, 并且空间复杂度是常数级
 */

func PCount(array []int) []int {
	retArray := make([]int, len(array))

	i, j, index := 0, 0, 0
	iLen := len(array)
	for i = 0; i < iLen; i++ {
		index = 0
		for j = 0; j < iLen; j++ {
			if array[i] > array[j] {
				index++
			}
		}
		retArray[index] = array[i]
	}
	return retArray
}
