package utils

import "math/rand"

func GetMaxVal(data []int) int {
	max := data[0]
	for _,val := range data {
		if max < val {
			max = val
		}
	}
	return max
}

func GetTestData() []int {
	data := make([]int, 100)

	for i:=0; i<len(data); i++ {
		num := rand.Intn(100)
		data[i] = num
	}
	return data
}


