package sort

/**************
** 1)趟数
** 2)相邻交换
** 3)不交换即中止
****************/
func Bubble_sort(array []uint32) {
	iLength := len(array)
	bSwitch := false

	for i := 0; i < iLength; i++ {
		/*相邻两个元素的比较*/
		for j := 1; j < iLength-i; j++ {
			if array[j-1] > array[j] {
				bSwitch = true
				array[j-1], array[j] = array[j], array[j-1]
			}
		}
		//不交换表示排序完成
		if !bSwitch {
			break
		}
		bSwitch = false
	}

}
