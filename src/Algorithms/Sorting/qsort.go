package Sorting

import (
	"math/rand"
)

//////////////////////////////////////////////////////////////////////
func merger_sort(data []int, begin int, ends int) {
	if begin >= ends {
		return
	}
	middle := (ends - begin) / 2  + begin
	merger_sort(data, begin, middle)
	merger_sort(data, middle+1, ends)
	merge(data, begin, middle, ends)
}

func merge(data []int ,begin int, middle int, ends int) {
	tmpt := make([]int, 0, ends - begin + 1)

	left := begin
	right := middle + 1
	for left <= middle && right <= ends {
		if data[left] < data[right] {
			tmpt = append(tmpt, data[left])
			left++
		} else {
			tmpt = append(tmpt,data[right])
			right++
		}
	}

	tmpt = append(tmpt, data[right:ends+1]...)
	tmpt = append(tmpt, data[left:middle+1]...)
	copy(data[begin: ends+1], tmpt)
}

//////////////////////////////////////////////////////////////////////
func qsort(data []int, begin int, end int) {
	if begin >= end {
		return
	}

	j := partition(data,begin,end)
	qsort(data,begin,j-1)
	qsort(data,j+1,end)
}

func Rand(begin int, end int) int {
	return begin + rand.Intn(end - begin + 1)
}

func swap(data []int, i int ,j int) {
	tmpt := data[i]
	data[i] = data[j]
	data[j] = tmpt
}

func partition(data []int, begin int, end int) int {
	swap(data, Rand(begin, end) , end)

	value := data[end]
	i := begin
	for j := begin; j <= end; j++ {
		if data[j] < value {
			if i != j {
				swap(data, i, j)
			}
			i++
		}
	}

	swap(data, i, end)
	return i
}