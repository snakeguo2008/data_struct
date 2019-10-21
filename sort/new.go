package sort

//import "fmt"

//合并两个数组
/*
1、极端参数校验
2、合并规则
* 对两个数组进行逐一比较，将较小的元素放入新数组切片中，知道任意一个数组结束；
* 随后将未结束的切片拼接到新切片尾部
*/
func MergeArray(a, b []int32) []int32 {

	if len(a) == 0 {
		return b
	}

	if len(b) == 0 {
		return a
	}

	var ret []int32

	i, j := 0, 0

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			ret = append(ret, a[i])
			i++
		} else {
			ret = append(ret, b[j])
			j++
		}
	}
	//下标越界会不会溢出
	ret = append(ret, a[i:]...)
	ret = append(ret, b[j:]...)

	return ret
}

//归并排序法，单线程
/*
*1、表层递归
*2、底层合并
 */

func MergerSort(array []int32, start, end int) []int32 {
	/*考虑递归  */
	//退出条件
	if start >= end {
		return array[start : end+1]
	}

	//递归分割
	mid := start + (end-start)/2
	left := MergerSort(array, start, mid)
	right := MergerSort(array, mid+1, end)
	/*考虑递归*/

	//考虑进入最后2层
	return MergeArray(left, right)
}

/*
*多线程合并靠考虑要点：
*1、父协程需要等子协程执行完，也就是设计到同步过程
*2、父协程需要获取子协程的合并结果,也就是设计到线程间通信
 */
func MultPSort(array []int32, start, end int, ret chan []int32) {
	/**考虑递归**/
	//递归出口
	if start >= end {
		ret <- array[start : end+1]
		return
	}

	//分割
	retCh := make(chan []int32, 2)
	mid := start + (end-start)/2
	go MultPSort(array, start, mid, retCh)
	go MultPSort(array, mid+1, end, retCh)

	retArr1 := <-retCh
	retArr2 := <-retCh
	close(retCh)

	/**合并***/
	data := MergeArray(retArr1, retArr2)

	ret <- data
}

/*
*快速排序分割子序列要保证父协程已经排完了才行进行自序列的分割，也就是设计到线程同步
*快速排序不需要申请多余的空间
*快排序分为两个过程，1、为一个数组找中值；2、分割
 */

func OrginMidMun(array []int32, start, end int) int {
	if start >= end {
		return -1
	}

	i, j := start, end
	//保存索引和值
	key := array[start]
	index := start

	for i <= j {
		for i <= j && array[j] >= key {
			j--
		}

		if i <= j {
			array[index] = array[j]
			index = j
		}

		for i <= j && array[i] <= key {
			i++
		}

		if i <= j {
			array[index] = array[i]
			index = i
		}
	}

	array[index] = key

	return index
}

/*
1、将快排拆分为递归出口和递归表达式两部分
2、先获取中值索引下标，然后再进行递归
*/
func MultiQSort(array []int32, start, end int, ch chan int) {
	//

	defer func() {
		ch <- 1
	}()

	if start >= end {

		return
	}
	//	获取中值
	index := OrginMidMun(array, start, end)

	chR := make(chan int, 1)

	//递归
	//if start < index-1 {
	go MultiQSort(array, start, index-1, chR)
	//}

	//if index+1 < end {
	go MultiQSort(array, index+1, end, chR)
	//}
	<-chR
	<-chR

	close(chR)

}
