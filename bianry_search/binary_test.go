package bianry_search_test

import (
	"testing"
)

func binary(key int, list []int) int {
	start, end := 0, len(list)-1
	for start <= end {
		//mid := start + (end-start)/2
		//mid := end + (start-end)/2
		mid := start + (end-start)/2
		if key > list[mid] {
			start = mid + 1
		} else if key < list[mid] {
			end = mid - 1
		} else {
			return mid
		}
	}
	return -1
}
func TestBinary(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7}
	if ret := binary(4, list); ret != 3 {
		t.Errorf("1.value: %d", ret)
	}
	if ret := binary(1, list); ret != 0 {
		t.Errorf("2.value: %d", ret)
	}
	if ret := binary(7, list); ret != 6 {
		t.Errorf("3.value: %d", ret)
	}
	if ret := binary(-11, list); ret != -1 {
		t.Errorf("4.value: %d", ret)
	}
	if ret := binary(422, list); ret != -1 {
		t.Errorf("5.value: %d", ret)
	}
}

func BinaryRecursion(key int, list []int, start, end int) int {
	if start <= end {
		mid := start + (end-start)/2
		if key > list[mid] {
			return BinaryRecursion(key, list, mid+1, end)
		} else if key < list[mid] {
			return BinaryRecursion(key, list, start, mid-1)
		} else {
			return mid
		}
	}
	return -1
}
func TestBinaryRecursion(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7}
	if ret := BinaryRecursion(4, list, 0, len(list)-1); ret != 3 {
		t.Errorf("1.value: %d", ret)
	}
	if ret := BinaryRecursion(1, list, 0, len(list)-1); ret != 0 {
		t.Errorf("1.value: %d", ret)
	}
	if ret := BinaryRecursion(-4, list, 0, len(list)-1); ret != -1 {
		t.Errorf("1.value: %d", ret)
	}
	if ret := BinaryRecursion(43, list, 0, len(list)-1); ret != -1 {
		t.Errorf("1.value: %d", ret)
	}

}
