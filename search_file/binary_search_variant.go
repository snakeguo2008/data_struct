package search_file

/*********二分查找变体********/

/**************
***查找第一个值等于给定值的元素
**************/
/*****当找到与目标元素相同的值时，如何判断他是否是第一个元素呢？
**当该元素前不存在别的元素，或者有别的元素但该该元素不等于目标元素时,那么他就是这个第一个匹配的元素;
**相反,相反那么第一元素必然在low----(mid-1)这个元素中间
 */

func Binary_search_first(array []int, x int) int {
	low := 0
	high := len(array) - 1

	for low <= high {
		mid := low + (high-low)>>1
		if x > array[mid] {
			low = mid + 1
		} else if x < array[mid] {
			high = mid - 1
		} else {
			//判断该元素之前是否存在相同的元素;需要注意mid=0的情况
			if (mid == 0) || array[mid-1] != x {
				return mid
			} else {
				high = mid - 1 //否则第一个元素必然在[low, mid-1]之前
			}
		}
	}
	return -1
}

/***************查找最后一个值等于给定值的元素*********
*** 类似的，当该元素为最后一个元素或者该元素之后的元素不等于目标元素时，即为我们想要找的位置
***
 */
func Binary_search_end(array []int, x int) int {
	low := 0
	high := len(array) - 1
	mid := 0

	for low <= high {
		mid = low + (high-low)>>1
		if x > array[mid] {
			low = mid + 1
		} else if x < array[mid] {
			high = mid - 1
		} else {
			//当钙元素时最后一个元素或者该元素后面的元素!=目标元素时，即为所找
			//否则最后一个元素则在[mid+1:high]之间
			if mid == len(array)-1 || array[mid+1] != x {
				return mid
			} else {
				low = mid + 1
			}

		}
	}
	return -1

}

/**********************
*** 变种3: 查找第一个大于等于给定值的元素
***
 */
func Binary_search_be(array []int, x int) int {
	low := 0
	high := len(array) - 1
	mid := 0

	for low <= high {
		mid = low + (high-low)>>1
		if x > array[mid] {
			low = mid + 1
		} else {
			//类似的，如果mid之前没有数据或者mid的数据小于x,则返回mid;否则要找的数据在[low:mid-1]
			if mid == 0 || array[mid-1] < x {
				return mid
			} else {
				high = mid - 1
			}

		}
	}
	return -1

}

/************************
*** 变种4: 查找最后一个小于等于给定值的元素
*** 3，5，6，8，9，10。最后一个小于等于 7 的元素就是 6
 */

func Binary_search_se(array []int, x int) int {
	low := 0
	high := len(array) - 1
	mid := 0

	for low <= high {
		mid = low + (high-low)>>1
		if x < array[mid] {
			high = mid - 1
		} else {
			//如果当前元素已经是最后一个元素或者该元素后面的元素已经比x大,那么该元素下标即为目标
			if mid == len(array)-1 || array[mid+1] > x {
				return mid
			} else {
				low = mid + 1
			}
		}
	}
	return -1
}
