package sort

/****************
*** 1、自底向上合并(即递归的顺序是先拆分再合并，和快排相反)
***
***
****************/

func MergeArr(arrayA, arrayB []uint32) []uint32 {
	if len(arrayA) == 0 {
		return arrayB
	}
	if len(arrayB) == 0 {
		return arrayA
	}

	iLengthA := len(arrayA)
	iLengthB := len(arrayB)
	i, j := 0, 0
	var retArray []uint32
	for i < iLengthA && j < iLengthB {
		if arrayA[i] > arrayB[j] {
			retArray = append(retArray, arrayB[j])
			j++
		} else {
			retArray = append(retArray, arrayA[i])
			i++
		}
	}

	retArray = append(retArray, arrayA[i:]...)
	retArray = append(retArray, arrayB[j:]...)

	return retArray
}

//递归方法，注意递归出口;
func MergeSort(array []uint32) []uint32 {
	if len(array) <= 1 {
		return array
	}

	iMid := len(array) / 2
	//由于是完全拆分，所以这里可以使用0为下标
	lArray := MergeSort(array[0:iMid])
	rArray := MergeSort(array[iMid:])

	return MergeArr(lArray, rArray)

}
