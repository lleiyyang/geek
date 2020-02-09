package Sorting

import "fmt"

type  MaxHeap struct{
	data []int
	length int
}

func NewMaxHeap(data []int) *MaxHeap {
	maxindex := len(data)
	heap := &MaxHeap{
		data: make([]int, maxindex + 1),
		length: maxindex + 1,
	}
	copy(heap.data[1:], data)

	for i := maxindex / 2; i >= 1; i-- {
		heap.Adjust_Down(i)
		heap.Print_Heap()
	}
	return heap
}

func (heap *MaxHeap) Print_Heap() {
	fmt.Println()
	data := heap.data
	for _, num := range data{
		fmt.Printf("%v ", num)
	}
}

func (heap *MaxHeap) Heap_Sort() {
	for heap.length > 1 {
		swap(heap.data, 1, heap.length-1)
		heap.length--
		heap.Adjust_Down(1)
	}
}

func (heap *MaxHeap) Adjust_Down(index int) {
	data := heap.data
	maxindex := heap.length - 1
	for index < maxindex {
		left_index := index * 2
		if left_index > maxindex {//无左节点
			return
		}
		if index * 2 + 1 > maxindex {//无右节点
			if data[index] >= data[left_index] {
				return
			}
			swap(data, index, left_index)
			index = left_index
			continue
		}

		//左节点即较大的节点
		if data[left_index] < data[index * 2 + 1] {
			left_index = index * 2 + 1
		}
		if data[index] >= data[left_index] {
			return
		}
		swap(data, index, left_index)
		index = left_index
	}
	return
}
