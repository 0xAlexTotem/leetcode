package main

func countNicePairs(nums []int) int {
	const mod = 1_000_000_007
	m := make(map[int]int)
	var count int

	for i := range nums {
		m[nums[i]-rev(nums[i])]++
	}

	for _, n := range m {
		count = (count + n*(n-1)/2) % mod
	}

	return count
}

func rev(n int) int {
	var r int

	for n > 0 {
		r = r*10 + n%10
		n /= 10
	}

	return r
}
