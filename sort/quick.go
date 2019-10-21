package sort

/*****************
* 1、分治，自顶向下的排序
* 2、注意递归出口
* 3、最坏时间复杂度分析，由于比较的次数都是n，不能减少。所以当递归的深度最大时(为n)，则
*    时间复杂度为O(n^2), 当刚好每次5/5分割时，递归深度为logn,所以最好的时间复杂度就是O(nlogn)
***/

//控制递归出口
func Quick_sort(left, right int, array []uint32) {
	if left >= right {
		return
	}

	iMid := quick_sort_1(left, right, array)
	//fmt.Printf("快排中间数:%v \n",iMid)

	if left < iMid-1 {
		Quick_sort(left, iMid-1, array)
	}

	if iMid+1 < right {
		Quick_sort(iMid+1, right, array)
	}

}

//这种并非正统的快速排序方法
func quick_merge(left, right int, array []uint32) int {

	i := left
	j := right
	iNum := array[left]

	for i < j {
		//找比iNum小的数
		for ; i < j && array[j] > iNum; j-- {

		}
		//找比iNum大的数
		for ; i < j && array[i] < iNum; i++ {

		}

		//交换这两个数
		if i < j {
			array[i], array[j] = array[j], array[i]
		}

	}

	array[i] = iNum

	return i
}

//挖坑填数正统快排法
func quick_sort_1(left, right int, array []uint32) int {

	i := left
	j := right
	iMidNum := array[left]
	index := i
	//**这里全是< or >号,没有等于号
	/**
	** 1处==>表示当i >=j代表结束
	** 2处==>表示即使和基准数相等,也会交换。这种实现体现了快排的不稳定排序特性
	**/
	for i < j { //1
		for ; i < j && array[j] > iMidNum; j-- { //2
		}

		if i < j {
			array[index] = array[j]
			index = j
		}

		for ; i < j && array[i] < iMidNum; i++ {
		}

		if i < j {
			array[index] = array[i]
			index = i
		}

	}

	array[index] = iMidNum

	return index

}
