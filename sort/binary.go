package sort

/*********
** 1、折半插入排序
** 2、是简单插入排序的升级版本，在查找插入位置时会比前者快,但是移动数据位置相同
** 3、时间复杂度,最好O(nlogn),最坏O(n^2) ;空间复杂度 O(1)
** 4、和简单插入排序一样都是下标为0的元素做标杆，所以i都从1开始
 */

func BinaryInsert(array []uint32) {
	temp := uint32(0)
	iLength := len(array)
	i, j, low, high := 0, 0, 0, 0

	for i = 1; i < iLength; i++ {
		temp = array[i]
		//这两个数值
		low = 0
		high = i - 1

		for low <= high {
			iMid := (low + high) / 2
			if temp > array[iMid] {
				low = iMid + 1

			} else {
				high = iMid - 1
			}
		}

		//移动元素位置

		for j = i - 1; j > high; j-- {

			array[j+1] = array[j]

		}

		array[j+1] = temp
	}

}
