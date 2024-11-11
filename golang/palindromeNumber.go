package main

import "strconv"

func isPalindrome(x int) bool {
	if x == 0 {
		return true
	}

	if x < 0 || x%10 == 0 {
		return false
	}

	s := strconv.Itoa(x)

	for l := 0; l < (len(s) / 2); l++ {
		if s[l] != s[len(s)-l-1] {
			return false
		}
	}

	return true
}
