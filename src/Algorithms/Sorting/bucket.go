package Sorting

import(
	"Algorithms/utils"
)

//桶排序
//1：初始化桶
//2：入桶：根据数值找到对应的桶
//3：桶内排序
//4：桶间排序

const (
	Bucket_Num int = 100
)

func Bucket_Sort(data []int) {
	buckets := make([][]int, Bucket_Num)

	//入桶
	max := utils.GetMaxVal(data)
	for _,val := range data {
		index := val * (Bucket_Num - 1) / max
		buckets[index] = append(buckets[index], val)
	}

	//桶内排序
	offset := 0
	for index,_ := range buckets {
		merger_sort(buckets[index], 0, len(buckets[index])-1)
		if len(buckets[index]) > 0 {
			copy(data[offset:], buckets[index])
			offset = offset + len(buckets[index])
		}
	}
}