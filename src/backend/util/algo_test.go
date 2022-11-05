package util

import "testing"

func TestFindBinarySearch(t *testing.T) {
	arr := []int{1,2,3,4,5,6,7,8,9,11,12,12,34}
	res,ok := BinarySearch(arr,7)
	t.Log(res,ok)
}