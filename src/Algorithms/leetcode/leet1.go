package leetcode

import (
	"fmt"
	"sort"
)

type myint []int

func (data myint) Len() int {
	return len(data)
}

func (data myint) Less(i, j int) bool {
	return data[i] > data[j]
}

func (data myint) Swap(i, j int) {
	tmpt := data[i]
	data[i] = data[j]
	data[j] = tmpt
}

func leastInterval(tasks []byte, n int) int {
	countmap := make(myint, 26)
	for _, task := range tasks {
		countmap[task - 'A']++
	}

	cnt := 0
	sort.Sort(countmap)
	fmt.Println(countmap)
	for countmap[0] > 0 {
		i := 0
		for i = 0; i <= n; i++ {
			if countmap[i] > 0 {
				countmap[i]--
				cnt++
			} else if countmap[0] > 0 {
				cnt = cnt + (n - i + 1)
				break
			} else {
				break
			}

		}
		sort.Sort(countmap)
	}
	return cnt
}