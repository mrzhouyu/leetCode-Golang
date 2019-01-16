package main

func twoSum(nums []int, target int) []int {
	var res []int
	for index, value := range nums[:len(nums)-1] {
		for index1, value1 := range nums[index+1:] {
			if value+value1 == target {
				res = append(res, index, index+index1+1)
				return res
			}
		}
	}
	return res
}
