package sort

/*****
**简单选择排序
** 1、取第一个数为基数,进行偏移操作
**
**/

func InsertArray(array []uint32) {
	iLength := len(array)
	i, j := 1, 0
	temp := uint32(0)
	for ; i < iLength; i++ {
		temp = array[i]

		for j = i - 1; j >= 0; j-- {
			if array[j] >= temp {
				array[j+1] = array[j]
			}
		}
		array[j+1] = temp
	}
}
