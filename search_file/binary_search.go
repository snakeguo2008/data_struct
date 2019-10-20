package search_file

import "fmt"

/********
***二分查找(数据从大到小排序;不存在重复元素)
***要点:
***
*** 1、当使用循环方式实现时 low <= high
*** 2、获取两者之间的中值时 low + (high-low)/2 不然容易溢出
*** 3、可以使用递归实现
***
 ********/

/*******************二分查找局限性**********
***  1、二分查找只能应用于顺序存储的数据
***  2、必须是排序的数据
***  3、数据量太小或者太大都不需要应用二分查找。太小遍历就足够了;
*** 	太大则因为需要为存储的数据分配连续的空间.
 */

/*******************
 *** 时间福复杂度分析:logn
 ***
 ***
 *********************/

/***循环实现****/
func Binary_search(array []int, target int) int {
	low := 0
	high := len(array) - 1
	mid := 0

	//这里需要理解循环退出条件是low > high;当low和high等值时的数据也可能是需要找的值。
	//例如只有一个数据且这个数据刚好是需要找的值
	for low <= high {
		mid = low + (high-low)/2
		if target > array[mid] {
			fmt.Println("bigger ", array[mid])
			low = mid + 1
		} else if target < array[mid] {
			fmt.Println("smaller ", array[mid])
			high = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

//递归实现二分查找
/***************
**递归要素:递归出口+递归表达式
**那么二分查找的递归出口是: low > high
 */
func Binary_search_cycle(low, high int, array []int, x int) int {
	if low > high {
		return -1
	}
	mid := low + (high-low)>>1
	if x > array[mid] {
		return Binary_search_cycle(mid+1, high, array, x)
	} else if x < array[mid] {
		return Binary_search_cycle(low, mid-1, array, x)
	} else {
		return mid
	}
	return -1
}
