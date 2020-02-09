package datastructure

import "sort"

func twoSum(nums []int, target int) []int {
	mapval := make(map[int]int)
	for i, num := range nums {
		mapval[i] = num
	}
	for i, num := range nums {
		if j, ok := mapval[target - num]; ok && i!= j {
			return []int{i,j}
		}
	}
	return []int{}
}

func lengthOfLongestSubstring(s string) int {
	max := 0
	right := 0
	lastleft := 0
	mapval := make(map[uint8]int)

	for left := 0; left <= len(s)-1; {
		for i := lastleft; i < left; i++{
			delete(mapval, s[i])
		}

		lastleft = left
		num := right - left
		j := right
		for j = right; j <= len(s)-1; j++ {
			 index ,ok := mapval[s[j]]
			 if !ok {
				mapval[s[j]] = j
				num++
			} else {
				right = j
				left = index + 1
				break
			}
		}
		if max < num {
			max = num
		}
		if j >= len(s) {
			return max
		}
	}
	return max
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	left := 0
	right := len(nums) - 1
	result := make([][]int, 0)

	for i := 0; i <= len(nums) - 1; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left = i + 1
		right = len(nums) - 1

		for left < right {
			sum := nums[left] + nums[right] + nums[i]
			if sum == 0 {
				result = append(result,[]int{nums[i], nums[left], nums[right]})
				j := left + 1
				for ; j < right; j++ {
					if nums[j] != nums[left] {
						break
					}
				}
				left = j
				right--
			} else if sum > 0 {
				right--
			}else {
				left++
			}
		}
	}

	return result
}