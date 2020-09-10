package main

/*
Given an array nums. We define a running sum of an array as runningSum[i] = sum(nums[0]…nums[i]).

Return the running sum of nums.



Example 1:

Input: nums = [1,2,3,4]
Output: [1,3,6,10]
Explanation: Running sum is obtained as follows: [1, 1+2, 1+2+3, 1+2+3+4].
Example 2:

Input: nums = [1,1,1,1,1]
Output: [1,2,3,4,5]
Explanation: Running sum is obtained as follows: [1, 1+1, 1+1+1, 1+1+1+1, 1+1+1+1+1].
*/

import "fmt"

func runningSum(nums []int) []int {
	var result []int
	preSum := 0
	for _, v := range nums {
		preSum += v
		result = append(result, preSum)
	}
	return result
}

//这种更节省内存
func runningSum1(nums []int) []int {
	for i := range nums {
		if i == 0 {
			continue
		}
		nums[i] = nums[i-1] + nums[i]
	}
	return nums
}

func main() {
	a := []int{3, 1, 2, 10, 1}
	fmt.Println(runningSum(a))
}
