package datastructure

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	str := "bbs"
	num := lengthOfLongestSubstring(str)
	t.Logf("str:%v, num:%v", str, num)
}


func Test_threeSum(t *testing.T) {
	nums := []int{-1,-1,0,1,2,-4}
	result := threeSum(nums)
	t.Logf("src:%v, result:%v", nums ,result)
}