package Sorting

import (
	"Algorithms/utils"
)

//计数排序： 桶排序的特殊情况，每个桶内只有一个元素，所以省掉了桶内排序的步骤
//算法流程：1：以算法的最大值为桶的大小 2：以桶的下标为具体数值, 入桶  3：为了保证稳定性，需要额外操作
//可以用来去重

func countingSort(data []int) {
	max := utils.GetMaxVal(data) + 1

	counts := make([]int, max)
	for _,val := range data {
		counts[val]++
	}


	//稳定版
	sum := 0
	for index, num := range counts {
		sum = sum + num
		counts[index] = sum
	}

	r := make([]int, len(data))
	for i := len(data) - 1; i >= 0; i-- {
		r[counts[data[i]] - 1] = data[i]
		counts[data[i]]--
	}
	copy(data, r)

	//不去重版本
	/*for index, val := range counts {
		for val > 0 {
			fmt.Printf("%v ",index)
			val--
		}
	}
	fmt.Println()*/

	//去重输出版本
	/*for index, val := range counts {
		if val > 0 {
			fmt.Printf("%v ",index)
		}
	}
	fmt.Println()*/
}