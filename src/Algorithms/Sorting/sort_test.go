package Sorting

import (
	"fmt"
	"testing"
	"Algorithms/utils"
)

func Test_Merger_Sort(t *testing.T){
	data := utils.GetTestData()
	merger_sort(data,0, len(data) - 1)
	if !test_order(data) {
		t.Errorf("sort failed. data:%v", data)
	}
}

func Test_Qsort(t *testing.T) {
	data := utils.GetTestData()
	qsort(data, 0, len(data) - 1)
	if !test_order(data) {
		t.Errorf("sort failed. data:%v", data)
	}
}

func Test_Bucket_Sort(t *testing.T) {
	data := utils.GetTestData()
	Bucket_Sort(data)
	if !test_order(data) {
		t.Errorf("sort failed. data:%v", data)
	}
}

func Test_countingSort(t *testing.T) {
	data := utils.GetTestData()
	countingSort(data)
	if !test_order(data) {
		t.Errorf("sort failed. data:%v", data)
	}
}

func Test_Insert_sort(t *testing.T) {
	data := utils.GetTestData()
	Insert_sort(data)
	if !test_order(data) {
		t.Errorf("sort failed. data:%v", data)
	}
}

func Test_shell_sort(t *testing.T) {
	data := utils.GetTestData()
	shell_sort(data)
	if !test_order(data) {
		t.Errorf("sort failed. data:%v", data)
	}
}

func Test_b_search(t *testing.T) {
	data := utils.GetTestData()
	shell_sort(data)
	index := b_search2(data,0, len(data)-1, data[100])
	if data[index] != data[100] {
		t.Error("b_search failed")
	}
}

func Test_mySqrt(t *testing.T) {
	val := 1000
	//ff := mySqrt(val)
	t.Logf("ff:%v", val)
}

func test_order(data []int) bool {
	end := len(data) - 1
	for i := 0; i < end; i++ {
		if data[i+1] < data[i] {
			return false
		}
	}

	return true
}

func Test_HeapSort(t *testing.T){
	data := utils.GetTestData()
	heap := NewMaxHeap(data)
	heap.Heap_Sort()
	fmt.Println("sort result:")
	heap.Print_Heap()
}
