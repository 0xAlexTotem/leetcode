package main

func twoSum(nums []int, target int) []int {
	for i, n := range nums {
		for k, m := range nums {
			if n == m && i == k {
				continue
			}

			if target == (n + m) {
				answer := []int{i, k}
				return answer
			}
		}
	}

	return []int{}
}
