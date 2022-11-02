package util

func BinarySearch(arr []uint64, targetValue uint64) (int, bool) {
	leftIndex := 0 // 左下标
	rightIndex := len(arr) - 1 // 右下标

	// 若leftIndex大于rightIndex，表示数组中没有该数据
	for leftIndex <= rightIndex {
		// 中间下标
		middleIndex := (leftIndex + rightIndex) / 2

		if arr[middleIndex] > targetValue {
			// 目标数据在左半边
			rightIndex = middleIndex - 1
		} else if arr[middleIndex] < targetValue {
			// 目标数据在右半边
			leftIndex = middleIndex + 1
		} else {
			// 找到目标数据
			return middleIndex, true
		}
	}
	return -1, false
}



func ISort(arr *[]uint64) {
	var tmp uint64
	count := len(*arr)
	for i := 1; i < count; i++ {
		for j := i; j > 0 ; j-- {
			if (*arr)[j-1] > (*arr)[j] {
				tmp = (*arr)[j-1]
				(*arr)[j-1] = (*arr)[j]
				(*arr)[j] = tmp
			}else {
				break
			}
		}
	}
}



func QuickSort(arr []uint64, left, right int)  {
	if left < right{
		pivot := arr[left]
		j := left
		for i := left; i < right; i++ {
			if arr[i] < pivot {
				j++
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
		arr[left], arr[j] = arr[j], arr[left]
		QuickSort(arr, left, j)
		QuickSort(arr, j+1, right)
	}
}