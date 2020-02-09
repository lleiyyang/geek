package Sorting

//data保证有序
//非递归版本
//最简单的二分查找， 没有重复数据
//优势： 近似查找有优势， 因为近似查找二叉树不好做
func b_search(data []int, val int) int {

	left := 0
	right := len(data) -1
	mid := 0
	for right >= left {
		mid = left + (right - left) >> 1
		if val == data[mid] {
			return mid
		}
		if val >  data[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

//递归版本

func b_search2(data []int, left int, right int, val int) int {
	if right < left {
		return -1
	}

	mid := left + (right - left) / 2
	index := mid
	if data[mid] == val {
		return index
	}

	if val > data[mid] {
		index = b_search2(data, mid + 1, right, val)
	} else {
		index = b_search2(data, left, mid - 1, val)
	}

	return index
}

//二分查找的变种，处理重复数据
//查找第一个等于val的下标
func b2_search(data []int, val int) int {
	left := 0
	right := len(data) -1
	mid := 0
	for right >= left {
		mid = left + (right - left) >> 1
		if val == data[mid] {
			if mid == 0 || data[mid - 1] != val {
				return mid
			}else {
				right = mid - 1
			}
		}
		if val >  data[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

//查找最后一个等于val的下标
func b3_search(data []int, val int) int {
	left := 0
	right := len(data) -1
	mid := 0
	for right >= left {
		mid = left + (right - left) >> 1
		if val == data[mid] {
			if mid == len(data)-1 || data[mid + 1] != val {
				return mid
			}else {
				left = mid + 1
			}
		}
		if val >  data[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

//查找第一个大于等于val的下标
func b4_search(data []int, val int) int {
	left := 0
	right := len(data) -1
	mid := 0
	for right > left {
		mid = left + (right - left) / 2
		if data[mid] >= val  {
			if mid == 0 || data[mid - 1] < val {
				return mid
			} else {
				right = mid - 1
			}
		} else {
			left = mid + 1
		}
	}
	return -1
}


//查找最后一个小于等于val的下标
func b5_search(data []int, val int) int {
	left := 0
	right := len(data) -1
	mid := 0
	for right > left {
		mid = left + (right - left) / 2
		if data[mid] <= val  {
			if  mid == len(data)-1 ||  data[mid + 1] > val {
				return mid
			}else {
				left = mid + 1
			}
		} else {
			right = mid - 1
		}
	}
	return -1
}